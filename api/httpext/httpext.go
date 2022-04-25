package httpext

import (
	"encoding/json"
	"fmt"
	"go-todo-api/api/logger"
	"net/http"
)

type Port int

func (p Port) Addr() string {
	return fmt.Sprintf(":%d", p)
}

func WriteJSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(data)
	if err != nil {
		logger.Error("error on encode write json: %s", err.Error())
	}
}

type JSONError struct {
	Error string `json:"error"`
}

func WriteJSONError(writer http.ResponseWriter, statusCode int, err error) {
	je := new(JSONError)
	if err != nil {
		je.Error = err.Error()
	}
	WriteJSON(writer, statusCode, je)
}
