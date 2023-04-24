package response

import (
	"encoding/json"
	"log"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

// function response
func NewResponse(status bool, message string, data interface{}) Response {
	return Response{status, message, nil, data}
}

// function response error
func NewErrorResponse(message string, err string, data interface{}) Response {
	return Response{false, message, err, data}
}

// function parse file json
func (r *Response) MustMarshal() []byte {
	j, err := json.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}

	return j
}
