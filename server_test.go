package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"encoding/json"
)

type Result struct {
	Status  string  `json:"status"`
	Message string `json:"message"`
}

func TestGETPlayers(t *testing.T) {
	keys := "cuy"
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	requestKey, _ := http.NewRequest(http.MethodGet, "/?key="+keys, nil)
	response := httptest.NewRecorder()

	halo(response, request)
	
	t.Run("get msg", func(t *testing.T) {
		got := response.Header().Get("Content-Type")
		b, _ := ioutil.ReadAll(response.Body)
		
		var res Result
		json.Unmarshal(b, &res)
		
		if got != "application/json" {
			t.Errorf("got %q, want %q", got, "application/json")
		}
		if res.Status != "Success" {
			t.Errorf("got %q, want %q", res.Status, "Success")
		}
		if res.Message != "Hello dunia" {
			t.Errorf("got %q, want %q", res.Message, "Hello dunia")
		}
	})
	
	halo(response, requestKey)
	t.Run("get msg", func(t *testing.T) {
		got := response.Header().Get("Content-Type")
		b, _ := ioutil.ReadAll(response.Body)
		
		var res Result
		json.Unmarshal(b, &res)
		
		if got != "application/json" {
			t.Errorf("got %q, want %q", got, "application/json")
		}
		if res.Status != "Success" {
			t.Errorf("got %q, want %q", res.Status, "Success")
		}
		if res.Message != "Hello "+keys {
			t.Errorf("got %q, want %q", res.Message, "Hello "+keys)
		}
	})
}