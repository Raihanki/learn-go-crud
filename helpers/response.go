package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseApi struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(writer http.ResponseWriter, code int, message string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	response := &ResponseApi{
		Message: message,
		Data:    data,
	}

	res, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(res)
}
