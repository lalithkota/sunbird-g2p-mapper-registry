package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/mapper-service/services"
	"github.com/sunbirdrc/mapper-service/swagger_gen/models"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

// SearchHandler handles a request for linking an entry
func SearchHandler() operations.PostG2pFamapSearchHandler {
	return &search{}
}

type search struct {
}

type OnSearchRequest struct {
	Header struct {
		models.MsgHeader
	} `json:"header"`
	Message struct {
		TransactionID  *models.TransactionID    `json:"transaction_id"`
		SearchResponse []*models.SearchResponse `json:"search_response"`
	} `json:"message"`
	Signature models.MsgSignature `json:"signature,omitempty"`
}

func (u search) Handle(params operations.PostG2pFamapSearchParams) middleware.Responder {
	var searchResponses []*models.SearchResponse
	registry := services.NewRegistry()
	searchResult := registry.SearchEntity(*params.Body.Message.SearchRequest.AttributeType, params.Body.Message.SearchRequest.AttributeValue)
	if len(searchResult) > 0 {
		for _, result := range searchResult {
			resolveResponse := u.createResolveResponse(params.Body.Message.SearchRequest, result)
			searchResponse := u.createSearchResponse(resolveResponse)
			searchResponses = append(searchResponses, searchResponse)
		}
	} else {
		searchResponse := u.createUnsuccessfulSearchResponse(params.Body.Message.SearchRequest, services.NewResultPayloadForFailure())
		searchResponses = append(searchResponses, searchResponse)
	}

	onUpdateRequest := u.createOnSearchRequestPayload(params, searchResponses)
	services.NewCallbackService(string(params.Body.Header.SenderURI)).Callback(onUpdateRequest, "on-search")

	response := operations.NewPostG2pFamapSearchDefault(200)
	response.Payload = &operations.PostG2pFamapSearchDefaultBody{
		Message: &operations.PostG2pFamapSearchDefaultBodyMessage{
			AckStatus: "ACK",
			Error:     nil,
			Timestamp: models.Timestamp(time.Now()),
		},
	}
	return response
}

func (u search) createOnSearchRequestPayload(params operations.PostG2pFamapSearchParams, searchResponses []*models.SearchResponse) OnSearchRequest {
	onSearchRequest := OnSearchRequest{
		Header: params.Body.Header,
		Message: struct {
			TransactionID  *models.TransactionID    `json:"transaction_id"`
			SearchResponse []*models.SearchResponse `json:"search_response"`
		}{
			TransactionID:  params.Body.Message.TransactionID,
			SearchResponse: searchResponses,
		},
		Signature: params.Body.Signature,
	}
	return onSearchRequest
}

func (u search) createUnsuccessfulSearchResponse(resolveRequest *models.SearchRequest, result services.Result) *models.SearchResponse {
	timestamp := models.Timestamp(time.Now())
	resolveResponse := &models.SearchResponse{
		Locale:              resolveRequest.Locale,
		Mapper:              nil,
		ReferenceID:         resolveRequest.ReferenceID,
		Status:              models.NewRequestStatus(FAILURE),
		StatusReasonCode:    "rjct.reference_id.invalid",
		StatusReasonMessage: result.Params.Errmsg,
		Timestamp:           &timestamp,
	}
	return resolveResponse
}

func (u search) createResolveResponse(resolveRequest *models.SearchRequest, result services.FinancialMapper) *models.ResolveResponse {
	timestamp := models.Timestamp(time.Now())
	resolveResponse := &models.ResolveResponse{
		AccountProviderInfo: &models.AccountProviderInfo{
			AdditionalInfo: "",
			Code:           toReference("GTBIRWRKXXX"),
			Name:           toReference("GT Bank Rawanda"),
			Subcode:        "bir",
		},
		AdditionalInfo:      &result.AdditionalInfo,
		Fa:                  models.FinancialAddress(result.Fa),
		ID:                  models.PersonID(result.Id),
		Locale:              resolveRequest.Locale,
		ReferenceID:         resolveRequest.ReferenceID,
		Status:              models.NewRequestStatus(SUCCESS),
		StatusReasonCode:    "",
		StatusReasonMessage: "",
		Timestamp:           &timestamp,
	}
	return resolveResponse
}

func (u search) createSearchResponse(resolveResponse *models.ResolveResponse) *models.SearchResponse {
	var resolveArray []*models.ResolveResponse
	resolveArray = append(resolveArray, resolveResponse)
	return &models.SearchResponse{
		Locale:              resolveResponse.Locale,
		Mapper:              resolveArray,
		ReferenceID:         resolveResponse.ReferenceID,
		Status:              resolveResponse.Status,
		StatusReasonCode:    "",
		StatusReasonMessage: "",
		Timestamp:           resolveResponse.Timestamp,
	}
}
