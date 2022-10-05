package main

import (
	"log"
	"net/http"
	//server "github.com/marcusquigley/gowithtests3/httpserver"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	store := &InMemoryPlayerStore{}
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
