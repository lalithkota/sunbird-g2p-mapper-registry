package handlers

import (
	"strings"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/mapper-service/services"
	"github.com/sunbirdrc/mapper-service/swagger_gen/models"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

// UpdateHandler handles a request for linking an entry
func UpdateHandler() operations.PutG2pMapperUpdateHandler {
	return &update{}
}

type update struct {
}

type OnUpdateRequest struct {
	Header struct {
		models.MsgHeader
	} `json:"header"`
	Message struct {
		TransactionID  *models.TransactionID    `json:"transaction_id"`
		UpdateResponse []*models.UpdateResponse `json:"update_response"`
	} `json:"message"`
	Signature models.MsgSignature `json:"signature,omitempty"`
}

func (u update) Handle(params operations.PutG2pMapperUpdateParams) middleware.Responder {
	var updateResponses []*models.UpdateResponse
	registry := services.NewRegistry()
	for _, updateRequest := range params.Body.Message.UpdateRequest {
		entityOsid := registry.SearchEntityById(string(*(updateRequest.ID)))
		if entityOsid != "" {
			updateEntityResponse := registry.UpdateEntity(*updateRequest, entityOsid)
			updateResponse := u.createUpdateResponse(updateRequest, updateEntityResponse)
			updateResponses = append(updateResponses, updateResponse)
		} else {
			updateResponse := u.createUpdateResponse(updateRequest, services.NewResultPayloadForFailure())
			updateResponses = append(updateResponses, updateResponse)
		}

	}

	onUpdateRequest := u.createOnUpdateRequestPayload(params, updateResponses)
	services.NewCallbackService(string(params.Body.Header.SenderURI)).Callback(onUpdateRequest, "on-update")

	response := operations.NewPutG2pMapperUpdateDefault(200)
	response.Payload = &operations.PutG2pMapperUpdateDefaultBody{
		Message: &operations.PutG2pMapperUpdateDefaultBodyMessage{
			AckStatus: "ACK",
			Error:     nil,
			Timestamp: models.Timestamp(time.Now()),
		},
	}
	return response
}

func (u update) createOnUpdateRequestPayload(params operations.PutG2pMapperUpdateParams, updateResponses []*models.UpdateResponse) OnUpdateRequest {
	onUpdateRequest := OnUpdateRequest{
		Header: params.Body.Header,
		Message: struct {
			TransactionID  *models.TransactionID    `json:"transaction_id"`
			UpdateResponse []*models.UpdateResponse `json:"update_response"`
		}{
			TransactionID:  params.Body.Message.TransactionID,
			UpdateResponse: updateResponses,
		},
		Signature: params.Body.Signature,
	}
	return onUpdateRequest
}

func (u update) createUpdateResponse(updateRequest *models.UpdateRequest, updateEntityResponse services.Result) *models.UpdateResponse {
	timestamp := models.Timestamp(time.Now())
	updateResponse := &models.UpdateResponse{
		AdditionalInfo:      updateRequest.AdditionalInfo,
		Fa:                  updateRequest.Fa,
		Locale:              updateRequest.Locale,
		ReferenceID:         updateRequest.ReferenceID,
		Status:              models.NewRequestStatus(SUCCESS),
		StatusReasonCode:    "",
		StatusReasonMessage: "",
		Timestamp:           &timestamp,
	}
	if updateEntityResponse.Params.Status == "UNSUCCESSFUL" {
		updateResponse.Status = models.NewRequestStatus(FAILURE)
		if strings.Contains(updateEntityResponse.Params.Errmsg, "duplicate") {
			updateResponse.StatusReasonCode = "rjct.reference_id.duplicate"
		} else {
			updateResponse.StatusReasonCode = "rjct.reference_id.invalid"
		}
		updateResponse.StatusReasonMessage = updateEntityResponse.Params.Errmsg
	}
	return updateResponse
}
