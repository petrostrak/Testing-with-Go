package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func newGetScorePlayer(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q but want %q", got, want)
	}
}

func TestGETPlayers(t *testing.T) {
	server := &PlayerServer{}

	t.Run("returns Petros score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Petros", nil)
		// net/http/httptest has a spy already made for us called ResponseRecorder so we can use that
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Maggie's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Maggie", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}
