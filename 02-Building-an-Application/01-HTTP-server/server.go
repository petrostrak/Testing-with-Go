package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.proccessWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}

}

func (p *PlayerServer) proccessWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)

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
