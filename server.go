package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	server := &http.Server{
		Addr: ":1234",
	}


	http.HandleFunc("/", halo)

	fmt.Println("Server started on port 1234")
	server.ListenAndServe()
}

func halo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
		Message: "Hello "+msg,
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
