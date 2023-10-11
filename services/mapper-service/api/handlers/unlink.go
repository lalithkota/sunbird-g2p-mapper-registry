package handlers

import (
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/mapper-service/services"
	"github.com/sunbirdrc/mapper-service/swagger_gen/models"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

// UnlinkHandler handles a request for linking an entry
func UnlinkHandler() operations.PostG2pMapperUnlinkHandler {
	return &unlink{}
}

type unlink struct {
}

type OnUnlinkRequest struct {
	Header struct {
		models.MsgHeader
	} `json:"header"`
	Message struct {
		TransactionID  *models.TransactionID    `json:"transaction_id"`
		UnlinkResponse []*models.UnlinkResponse `json:"unlink_response"`
	} `json:"message"`
	Signature models.MsgSignature `json:"signature,omitempty"`
}

func (u unlink) Handle(params operations.PostG2pMapperUnlinkParams) middleware.Responder {
	var unlinkResponses []*models.UnlinkResponse
	registry := services.NewRegistry()
	for _, unlinkRequest := range params.Body.Message.UnlinkRequest {
		entityOsid := registry.SearchEntityById(string(*(unlinkRequest.ID)))
		if entityOsid != "" {
			deleteEntityResponse := registry.DeleteEntity(entityOsid)
			unlinkResponse := u.createUnlinkResponse(unlinkRequest, deleteEntityResponse)
			unlinkResponses = append(unlinkResponses, unlinkResponse)
		} else {
			unlinkResponse := u.createUnlinkResponse(unlinkRequest, services.NewResultPayloadForFailure())
			unlinkResponses = append(unlinkResponses, unlinkResponse)
		}
	}

	onUpdateRequest := u.createOnUnlinkRequestPayload(params, unlinkResponses)
	services.NewCallbackService(string(params.Body.Header.SenderURI)).Callback(onUpdateRequest, "on-unlink")

	response := operations.NewPostG2pMapperUnlinkDefault(200)
	response.Payload = &operations.PostG2pMapperUnlinkDefaultBody{
		Message: &operations.PostG2pMapperUnlinkDefaultBodyMessage{
			AckStatus: "ACK",
			Error:     nil,
			Timestamp: models.Timestamp(time.Now()),
		},
	}
	return response
}

func (u unlink) createOnUnlinkRequestPayload(params operations.PostG2pMapperUnlinkParams, unlinkResponses []*models.UnlinkResponse) OnUnlinkRequest {
	onUnlinkRequest := OnUnlinkRequest{
		Header: params.Body.Header,
		Message: struct {
			TransactionID  *models.TransactionID    `json:"transaction_id"`
			UnlinkResponse []*models.UnlinkResponse `json:"unlink_response"`
		}{
			TransactionID:  params.Body.Message.TransactionID,
			UnlinkResponse: unlinkResponses,
		},
		Signature: params.Body.Signature,
	}
	return onUnlinkRequest
}

func (u unlink) createUnlinkResponse(unlinkRequest *models.UnlinkRequest, unlinkEntityResponse services.Result) *models.UnlinkResponse {
	timestamp := models.Timestamp(time.Now())
	unlinkResponse := &models.UnlinkResponse{
		AdditionalInfo:      unlinkRequest.AdditionalInfo,
		Locale:              unlinkRequest.Locale,
		ReferenceID:         unlinkRequest.ReferenceID,
		Status:              models.NewRequestStatus(SUCCESS),
		StatusReasonCode:    "",
		StatusReasonMessage: "",
		Timestamp:           &timestamp,
	}
	if unlinkEntityResponse.Params.Status == "UNSUCCESSFUL" {
		unlinkResponse.Status = models.NewRequestStatus(FAILURE)
		if strings.Contains(unlinkEntityResponse.Params.Errmsg, "duplicate") {
			unlinkResponse.StatusReasonCode = "rjct.reference_id.duplicate"
		} else {
			unlinkResponse.StatusReasonCode = "rjct.reference_id.invalid"
		}
		unlinkResponse.StatusReasonMessage = unlinkEntityResponse.Params.Errmsg
	}
	return unlinkResponse
}
