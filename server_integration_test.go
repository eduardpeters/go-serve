package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/eduardpeters/go-serve"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := poker.NewFileSystemPlayerStore(database)
	assertNoError(t, err)

	server, err := poker.NewPlayerServer(store, dummyGame)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
