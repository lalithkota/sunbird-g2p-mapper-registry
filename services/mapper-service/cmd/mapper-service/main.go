package main

import (
	"flag"
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"
	"github.com/sunbirdrc/mapper-service/api/handlers"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi"
	"github.com/sunbirdrc/mapper-service/swagger_gen/restapi/operations"
)

func main() {
	var portFlag = flag.Int("port", 3000, "Port to run this service on")
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewMapperServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle
	api.PostG2pMapperLinkHandler = handlers.LinkHandler()
	api.PutG2pMapperUpdateHandler = handlers.UpdateHandler()
	api.PostG2pMapperUnlinkHandler = handlers.UnlinkHandler()
	api.PostG2pMapperResolveHandler = handlers.ResolveHandler()
	api.PostG2pFamapSearchHandler = handlers.SearchHandler()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
