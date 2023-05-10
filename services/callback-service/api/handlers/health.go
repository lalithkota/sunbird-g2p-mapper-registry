package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/callback-service/swagger_gen/restapi/operations"
)

// HealthHandler handles a request for linking an entry
func HealthHandler() operations.GetHealthHandler {
	return &health{}
}

type health struct {
}

func (h health) Handle(params operations.GetHealthParams) middleware.Responder {
	ok := operations.NewGetHealthOK()
	ok.Payload.Status = "UP"
	return ok
}
