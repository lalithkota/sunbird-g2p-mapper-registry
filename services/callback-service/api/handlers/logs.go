package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sunbirdrc/callback-service/swagger_gen/restapi/operations"
)

var transactionLogs = map[string][]string{}

func appendLog(transactionId string, msg string) {
	_, ok := transactionLogs[transactionId]
	if ok {
		transactionLogs[transactionId] = append(transactionLogs[transactionId], msg)
	} else {
		transactionLogs[transactionId] = []string{msg}
	}
}

func getLog(transactionId string) []string {
	logs, ok := transactionLogs[transactionId]
	if ok {
		return logs
	} else {
		return nil
	}
}

// LogsHandler handles a request for linking an entry
func LogsHandler() operations.GetLogsHandler {
	return &logs{}
}

type logs struct {
}

func (h logs) Handle(params operations.GetLogsParams) middleware.Responder {
	ok := operations.NewGetLogsOK()
	ok.SetPayload(&operations.GetLogsOKBody{
		Logs: getLog(params.TransactionID),
	})
	return ok
}
