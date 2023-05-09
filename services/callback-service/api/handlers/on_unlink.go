package handlers

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/sunbirdrc/callback-service/swagger_gen/models"
	"github.com/sunbirdrc/callback-service/swagger_gen/restapi/operations"
	"time"
)

// OnUnlinkHandler handles a request for linking an entry
func OnUnlinkHandler() operations.PostG2pMapperOnUnlinkHandler {
	return &onUnlink{}
}

type onUnlink struct {
}

func (l onUnlink) Handle(params operations.PostG2pMapperOnUnlinkParams) middleware.Responder {
	bytes, _ := json.Marshal(params.Body)
	log.Infof("On link response: %v", string(bytes))
	response := operations.NewPostG2pMapperOnLinkDefault(200)
	response.Payload = &operations.PostG2pMapperOnLinkDefaultBody{
		Message: &operations.PostG2pMapperOnLinkDefaultBodyMessage{
			AckStatus: "ACK",
			Error:     nil,
			Timestamp: models.Timestamp(time.Now()),
		},
	}
	return response
}
