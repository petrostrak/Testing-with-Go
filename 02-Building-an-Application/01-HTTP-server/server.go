package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players")

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	if player == "Petros" {
		return "20"
	}

	if player == "Maggie" {
		return "10"
	}

	return ""
}
