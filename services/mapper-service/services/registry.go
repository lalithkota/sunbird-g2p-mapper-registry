package services

import (
	"fmt"
	"reflect"

	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"github.com/sunbirdrc/mapper-service/config"
	"github.com/sunbirdrc/mapper-service/swagger_gen/models"
)

const FinancialAddressMapper = "FinancialAddressMapper"

var registryClient = Registry{}

type Registry struct {
	*req.Client
}

type Result struct {
	Id      string `json:"id"`
	Version string `json:"ver"`
	Ets     int64  `json:"ets"`
	Params  struct {
		Resmsgid string `json:"resmsgid"`
		Msgid    string `json:"msgid"`
		Err      string `json:"err"`
		Status   string `json:"status"`
		Errmsg   string `json:"errmsg"`
	} `json:"params"`
	ResponseCode string `json:"responseCode"`
	Result       struct {
		FinancialAddressMapper struct {
			Osid           string              `json:"osid,omitempty"`
			AdditionalInfo models.KeyValueInfo `json:"additional_info,omitempty"`

			// fa
			Fa models.FinancialAddress `json:"fa,omitempty"`

			// id
			ID models.PersonID `json:"id,omitempty"`

			// locale
			Locale models.LanguageCode `json:"locale,omitempty"`

			MobileNumber string `json:"mobile_number,omitempty"`

			// name
			Name models.LinkRequestName `json:"name,omitempty"`

			// reference id
			// Required: true
			ReferenceID models.ReferenceID `json:"reference_id"`

			// timestamp
			// Required: true
			// Format: date-time
			Timestamp models.Timestamp `json:"timestamp"`
		} `json:"FinancialAddressMapper"`
	} `json:"result"`
}

type equalFilter struct {
	Eq string `json:"eq"`
}

type filters struct {
	Filters struct {
		Id equalFilter `json:"id"`
	} `json:"filters"`
}

type FinancialMapper struct {
	Id             string                 `json:"id"`
	Fa             string                 `json:"fa"`
	Osid           string                 `json:"osid"`
	AdditionalInfo models.KeyValueInfo    `json:"additional_info,omitempty"`
	Locale         models.LanguageCode    `json:"locale,omitempty"`
	MobileNumber   string                 `json:"mobile_number,omitempty"`
	Name           models.LinkRequestName `json:"name,omitempty"`
	ReferenceID    models.ReferenceID     `json:"reference_id"`
	Timestamp      models.Timestamp       `json:"timestamp"`
}

func NewRegistry() Registry {
	if registryClient.Client == nil {
		c := req.C().
			// All GitHub API requests need this header.
			DevMode().
			SetCommonHeader("Accept", "application/json").
			// All GitHub API requests use the same base URL.
			SetBaseURL(config.Config.Registry.Url).
			// Enable dump at the request-level for each request, and only
			// temporarily stores the dump content in memory, so we can call
			// resp.Dump() to get the dump content when needed in response
			// middleware.
			// This is actually a syntax sugar, implemented internally using
			// request middleware
			EnableDumpEachRequest()
		registryClient = Registry{
			Client: c,
		}
	}
	return registryClient
}

func (r Registry) CreateEntity(request models.LinkRequest) Result {
	var result Result
	resp, err := r.R().
		SetBody(request).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		Post(fmt.Sprintf("/api/v1/%s", FinancialAddressMapper))
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
	}
	return result
}

func (r Registry) SearchEntityById(id string) string {
	var result []FinancialMapper
	var errResult Result
	filters := filters{
		Filters: struct {
			Id equalFilter `json:"id"`
		}{
			Id: equalFilter{id},
		},
	}
	resp, err := r.R().
		SetBody(filters).
		SetSuccessResult(&result).
		SetErrorResult(&errResult).
		Post(fmt.Sprintf("/api/v1/%s/search", FinancialAddressMapper))
	if err != nil {
		log.Fatal(err)
		log.Fatal(errResult)
	}

	if !resp.IsSuccessState() {
		log.Error("bad response status:", resp.Status)
	}
	if len(result) == 0 {
		log.Error("No mapper found", id)
		return ""
	}
	return result[0].Osid
}

func (r Registry) UpdateEntity(updateRequest models.UpdateRequest, osid string) Result {
	var result Result
	resp, err := r.R().
		SetBody(updateRequest).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		Put(fmt.Sprintf("/api/v1/%s/%s", FinancialAddressMapper, osid))
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
	}
	return result
}

func (r Registry) DeleteEntity(osid string) Result {
	var result Result
	resp, err := r.R().
		SetSuccessResult(&result).
		SetErrorResult(&result).
		Delete(fmt.Sprintf("/api/v1/%s/%s", FinancialAddressMapper, osid))
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
	}
	return result
}

func (r Registry) GetEntity(osid string) Result {
	var result Result
	resp, err := r.R().
		SetSuccessResult(&result).
		SetErrorResult(&result).
		Get(fmt.Sprintf("/api/v1/%s/%s", FinancialAddressMapper, osid))
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
	}
	return result
}

func NewResultPayloadForFailure() Result {
	return Result{
		Id:      "",
		Version: "",
		Ets:     0,
		Params: struct {
			Resmsgid string `json:"resmsgid"`
			Msgid    string `json:"msgid"`
			Err      string `json:"err"`
			Status   string `json:"status"`
			Errmsg   string `json:"errmsg"`
		}{
			Status: "UNSUCCESSFUL",
			Errmsg: "Invalid id",
		},
		ResponseCode: "",
		Result: struct {
			FinancialAddressMapper struct {
				Osid           string                  `json:"osid,omitempty"`
				AdditionalInfo models.KeyValueInfo     `json:"additional_info,omitempty"`
				Fa             models.FinancialAddress `json:"fa,omitempty"`
				ID             models.PersonID         `json:"id,omitempty"`
				Locale         models.LanguageCode     `json:"locale,omitempty"`
				MobileNumber   string                  `json:"mobile_number,omitempty"`
				Name           models.LinkRequestName  `json:"name,omitempty"`
				ReferenceID    models.ReferenceID      `json:"reference_id"`
				Timestamp      models.Timestamp        `json:"timestamp"`
			} `json:"FinancialAddressMapper"`
		}{},
	}
}

func (r Registry) SearchEntity(searchType string, value interface{}) []FinancialMapper {
	var result []FinancialMapper
	var errResult Result
	rt := reflect.TypeOf(value)
	var filters map[string]map[string]map[string]interface{}
	var operation = "eq"
	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {
		operation = "or"
	}
	filters = map[string]map[string]map[string]interface{}{
		"filters": {
			searchType: map[string]interface{}{
				operation: value,
			},
		},
	}
	resp, err := r.R().
		SetBody(filters).
		SetSuccessResult(&result).
		SetErrorResult(&errResult).
		Post(fmt.Sprintf("/api/v1/%s/search", FinancialAddressMapper))
	if err != nil {
		log.Fatal(err)
		log.Fatal(errResult)
	}

	if !resp.IsSuccessState() {
		log.Error("bad response status:", resp.Status)
	}
	return result
}
