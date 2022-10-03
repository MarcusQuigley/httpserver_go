package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	//server "github.com/marcusquigley/gowithtests3/httpserver"
)

func TestXxx(t *testing.T) {
	t.Run("return peppers score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("return Floyds score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		assertResponseBody(t, response.Body.String(), "200")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
