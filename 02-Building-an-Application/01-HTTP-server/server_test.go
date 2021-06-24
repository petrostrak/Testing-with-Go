package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Petros score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Petros", nil)
		// net/http/httptest has a spy already made for us called ResponseRecorder so we can use that
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})
}
