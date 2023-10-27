package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/mapper-service/services"
	"github.com/sunbirdrc/mapper-service/swagger_gen/models"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

// ResolveHandler handles a request for linking an entry
func ResolveHandler() operations.PostG2pMapperResolveHandler {
	return &resolve{}
}

type resolve struct {
}

type OnResolveRequest struct {
	Header struct {
		models.MsgHeader
	} `json:"header"`
	Message struct {
		TransactionID   *models.TransactionID     `json:"transaction_id"`
		ResolveResponse []*models.ResolveResponse `json:"resolve_response"`
	} `json:"message"`
	Signature models.MsgSignature `json:"signature,omitempty"`
}

func (u resolve) Handle(params operations.PostG2pMapperResolveParams) middleware.Responder {
	var resolveResponses []*models.ResolveResponse
	registry := services.NewRegistry()
	for _, resolveRequest := range params.Body.Message.ResolveRequest {
		var entities []services.FinancialMapper
		// If ID is given, always use ID to resolve
		// Else If FA is given, use FA to resolve
		if resolveRequest.ID != "" {
			entities = registry.SearchEntity("id", string(resolveRequest.ID))
		} else if resolveRequest.Fa != "" {
			// TODO: Since fa is not supposed to be unique in schema, multiple ids could be
			//   attached with same fa. id1 <-> fa1, id2 <-> fa1, id3 <-> fa1.
			//   Right now this only return the first id <-> fa mapping that it finds.
			//   Fix this issue.
			entities = registry.SearchEntity("fa", string(resolveRequest.Fa))
		}
		resolveResponse := u.createResolveResponse(resolveRequest, entities)
		resolveResponses = append(resolveResponses, resolveResponse)
	}

	onUpdateRequest := u.createOnResolveRequestPayload(params, resolveResponses)
	services.NewCallbackService(string(params.Body.Header.SenderURI)).Callback(onUpdateRequest, "on-resolve")

	response := operations.NewPostG2pMapperResolveDefault(200)
	response.Payload = &operations.PostG2pMapperResolveDefaultBody{
		Message: &operations.PostG2pMapperResolveDefaultBodyMessage{
			AckStatus: "ACK",
			Error:     nil,
			Timestamp: models.Timestamp(time.Now()),
		},
	}
	return response
}

func (u resolve) createOnResolveRequestPayload(params operations.PostG2pMapperResolveParams, resolveResponses []*models.ResolveResponse) OnResolveRequest {
	onResolveRequest := OnResolveRequest{
		Header: params.Body.Header,
		Message: struct {
			TransactionID   *models.TransactionID     `json:"transaction_id"`
			ResolveResponse []*models.ResolveResponse `json:"resolve_response"`
		}{
			TransactionID:   params.Body.Message.TransactionID,
			ResolveResponse: resolveResponses,
		},
		Signature: params.Body.Signature,
	}
	return onResolveRequest
}
func toReference(s string) *string {
	return &s
}

func (u resolve) createResolveResponse(resolveRequest *models.ResolveRequest, resolveEntityResponse []services.FinancialMapper) *models.ResolveResponse {
	timestamp := models.Timestamp(time.Now())
	resolveResponse := &models.ResolveResponse{
		AdditionalInfo:      resolveRequest.AdditionalInfo,
		Locale:              resolveRequest.Locale,
		ReferenceID:         resolveRequest.ReferenceID,
		Status:              models.NewRequestStatus(SUCCESS),
		StatusReasonCode:    "",
		StatusReasonMessage: "",
		Fa:                  "",
		ID:                  "",
		Timestamp:           &timestamp,
		AccountProviderInfo: nil,
	}
	if len(resolveEntityResponse) <= 0 {
		resolveResponse.Status = models.NewRequestStatus(FAILURE)
		// if ID is given return id.invalid
		// else if FA is given, return fa.invalid
		// if neither is given return id.invalid, because it expects id primarily.
		if resolveRequest.ID != "" {
			resolveResponse.StatusReasonCode = "rjct.id.invalid"
			resolveResponse.StatusReasonMessage = "ID is invalid or not found."
		} else if resolveRequest.Fa != "" {
			resolveResponse.StatusReasonCode = "rjct.fa.invalid"
			resolveResponse.StatusReasonMessage = "FA is invalid or not found."
		} else {
			resolveResponse.StatusReasonCode = "rjct.id.invalid"
			resolveResponse.StatusReasonMessage = "ID is invalid or not found."
		}
	} else {
		resolveResponse.ID = models.PersonID(resolveEntityResponse[0].Id)
		resolveResponse.Fa = models.FinancialAddress(resolveEntityResponse[0].Fa)
		resolveResponse.AccountProviderInfo = &models.AccountProviderInfo{
			AdditionalInfo: "",
			Code:           toReference("GTBIRWRKXXX"),
			Name:           toReference("GT Bank Rawanda"),
			Subcode:        "bir",
		}
	}
	return resolveResponse
}
