package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	//server "github.com/marcusquigley/gowithtests3/httpserver"
)

func TestXxx(t *testing.T) {
	t.Run("return peppers score", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
	t.Run("return Floyds score", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "200"
		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
