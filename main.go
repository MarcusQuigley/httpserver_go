package main

import (
	"log"
	"net/http"
	//server "github.com/marcusquigley/gowithtests3/httpserver"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
