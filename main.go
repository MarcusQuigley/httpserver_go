package main

import (
	"log"
	"net/http"
)

func main() {
	store := &InMemoryPlayerStore{}
	server := &PlayerServer{store: store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
