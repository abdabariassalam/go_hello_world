package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", halo)
	fmt.Println("Server started on port 1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}

const (
	StatusCtxKey                = 0
	StatusSuccess               = http.StatusOK
	StatusErrorForm             = http.StatusBadRequest
	StatusErrorUnknown          = http.StatusBadGateway
	StatusInternalError         = http.StatusInternalServerError
	StatusUnauthorized          = http.StatusUnauthorized
	StatusCreated               = http.StatusCreated
	StatusAccepted              = http.StatusAccepted
	StatusForbidden             = http.StatusForbidden
	StatusInvalidAuthentication = http.StatusProxyAuthRequired
)

var statusMap = map[int][]string{
	StatusSuccess:               {"STATUS_OK", "Success"},
	StatusErrorForm:             {"STATUS_BAD_REQUEST", "Invalid data request"},
	StatusErrorUnknown:          {"STATUS_BAG_GATEWAY", "Oops something went wrong"},
	StatusInternalError:         {"INTERNAL_SERVER_ERROR", "Oops something went wrong"},
	StatusUnauthorized:          {"STATUS_UNAUTHORIZED", "Not authorized to access the service"},
	StatusCreated:               {"STATUS_CREATED", "Resource has been created"},
	StatusAccepted:              {"STATUS_ACCEPTED", "Resource has been accepted"},
	StatusForbidden:             {"STATUS_FORBIDDEN", "Forbidden access the resource "},
	StatusInvalidAuthentication: {"STATUS_INVALID_AUTHENTICATION", "The resource owner or authorization server denied the request"},
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    string      `json:"code"`
}

func halo(w http.ResponseWriter, r *http.Request) Response {

	keys := r.URL.Query()["key"]
	msg := "dunia"
	if len(keys) > 0 {
		msg = (keys[0])
	}

	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{

		Status:  "Success",
		Message: "Hello " + msg,
	}

	response, err := json.Marshal(payload)
	if err != nil {
		return Response{}
	}

	return Response{
		Data:    response,
		Message: statusMap[http.StatusOK][1],
		Code:    statusMap[http.StatusOK][0],
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, heroku")
}
