package handlers

import (
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/mapper-service/services"
	"github.com/sunbirdrc/mapper-service/swagger_gen/models"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

const SUCCESS = "succ"
const FAILURE = "rjct"

// LinkHandler handles a request for linking an entry
func LinkHandler() operations.PostG2pMapperLinkHandler {
	return &link{}
}

type link struct {
}

type OnLinkRequest struct {
	Header struct {
		models.MsgHeader
	} `json:"header"`
	Message struct {
		TransactionID *models.TransactionID  `json:"transaction_id"`
		LinkResponse  []*models.LinkResponse `json:"link_response"`
	} `json:"message"`
	Signature models.MsgSignature `json:"signature,omitempty"`
}

func (l link) Handle(params operations.PostG2pMapperLinkParams) middleware.Responder {
	var linkResponses []*models.LinkResponse

	for _, linkRequest := range params.Body.Message.LinkRequest {
		createEntityResponse := services.NewRegistry().CreateEntity(*linkRequest)
		linkResponse := l.createLinkResponse(linkRequest, createEntityResponse)
		linkResponses = append(linkResponses, linkResponse)
	}

	onLinkRequest := l.createOnLinkRequestPayload(params, linkResponses)
	services.NewCallbackService(string(params.Body.Header.SenderURI)).Callback(onLinkRequest, "on-link")

	response := operations.NewPostG2pMapperLinkDefault(200)
	response.Payload = &operations.PostG2pMapperLinkDefaultBody{
		Message: &operations.PostG2pMapperLinkDefaultBodyMessage{
			AckStatus: "ACK",
			Error:     nil,
			Timestamp: models.Timestamp(time.Now()),
		},
	}
	return response
}

func (l link) createOnLinkRequestPayload(params operations.PostG2pMapperLinkParams, linkResponses []*models.LinkResponse) OnLinkRequest {
	onLinkRequest := OnLinkRequest{
		Header: params.Body.Header,
		Message: struct {
			TransactionID *models.TransactionID  `json:"transaction_id"`
			LinkResponse  []*models.LinkResponse `json:"link_response"`
		}{
			TransactionID: params.Body.Message.TransactionID,
			LinkResponse:  linkResponses,
		},
		Signature: params.Body.Signature,
	}
	return onLinkRequest
}

func (l link) createLinkResponse(linkRequest *models.LinkRequest, createEntityResponse services.Result) *models.LinkResponse {
	timestamp := models.Timestamp(time.Now())
	linkResponse := &models.LinkResponse{
		AdditionalInfo:      linkRequest.AdditionalInfo,
		Fa:                  *linkRequest.Fa,
		Locale:              linkRequest.Locale,
		ReferenceID:         linkRequest.ReferenceID,
		Status:              models.NewRequestStatus(SUCCESS),
		StatusReasonCode:    "",
		StatusReasonMessage: "",
		Timestamp:           &timestamp,
	}
	if createEntityResponse.Params.Status == "UNSUCCESSFUL" {
		linkResponse.Status = models.NewRequestStatus(FAILURE)
		if strings.Contains(createEntityResponse.Params.Errmsg, "duplicate") {
			linkResponse.StatusReasonCode = "rjct.reference_id.duplicate"
		} else {
			linkResponse.StatusReasonCode = "rjct.reference_id.invalid"
		}
		linkResponse.StatusReasonMessage = createEntityResponse.Params.Errmsg
	}
	return linkResponse
}
