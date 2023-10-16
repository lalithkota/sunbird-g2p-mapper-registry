package services

import (
	"fmt"

	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"github.com/sunbirdrc/mapper-service/config"
)

func NewCallbackService(callbackUrl string) CallbackService {
	var url string
	if callbackUrl != "" {
		url = callbackUrl
	} else {
		url = config.Config.CallbackService.Url
	}
	c := req.C().
		// All GitHub API requests need this header.
		DevMode().
		SetCommonHeader("Accept", "application/json").
		// All GitHub API requests use the same base URL.
		SetBaseURL(url).
		// Enable dump at the request-level for each request, and only
		// temporarily stores the dump content in memory, so we can call
		// resp.Dump() to get the dump content when needed in response
		// middleware.
		// This is actually a syntax sugar, implemented internally using
		// request middleware
		EnableDumpEachRequest()
	return CallbackService{
		Client: c,
	}
}

type CallbackService struct {
	*req.Client
}

func (r CallbackService) Callback(request interface{}, callbackMethod string) interface{} {
	var result interface{}
	resp, err := r.R().
		SetBody(request).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		Post(fmt.Sprintf("/%s", callbackMethod))
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccessState() {
		fmt.Println("bad response status:", resp.Status)
	}
	return result
}
