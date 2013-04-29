package resty

import (
	"net/http"
	"log"
	"fmt"
	"runtime/debug"
	"encoding/json"
)

type ErrorMessage struct {
	Message string
	Stack string
}

func RestError(err error, response http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		var errorMsg ErrorMessage
		errorMsg.Message = fmt.Sprintf("%v", err)
		errorMsg.Stack = string(debug.Stack())
		out, _ := json.Marshal(errorMsg)
		response.Header().Set("Content-Type", "application/json")
		response.Write(out)
		return true
	}
	return false
}
