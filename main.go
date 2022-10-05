package main

import (
	"log"
	"net/http"
	//server "github.com/marcusquigley/gowithtests3/httpserver"
)

func main() {
	store := &InMemoryPlayerStore{}
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
