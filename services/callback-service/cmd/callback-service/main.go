package main

import (
	"flag"
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"
	"github.com/sunbirdrc/callback-service/api/handlers"
	"github.com/sunbirdrc/callback-service/swagger_gen/restapi"
	"github.com/sunbirdrc/callback-service/swagger_gen/restapi/operations"
)

func main() {
	var portFlag = flag.Int("port", 3001, "Port to run this service on")
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewCallbackServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle

	api.PostG2pMapperOnLinkHandler = handlers.OnLinkHandler()
	api.PutG2pMapperOnUpdateHandler = handlers.OnUpdateHandler()
	api.PostG2pMapperOnUnlinkHandler = handlers.OnUnlinkHandler()
	api.PostG2pMapperOnResolveHandler = handlers.OnResolveHandler()
	api.PostG2pFamapOnSearchHandler = handlers.OnSearchHandler()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
