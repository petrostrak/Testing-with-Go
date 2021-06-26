package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8080", server))
}
