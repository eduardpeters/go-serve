package poker_test

import (
	"io"
	"testing"
	"time"

	poker "github.com/eduardpeters/go-serve"
)

var dummyPlayerStore = &poker.StubPlayerStore{}

func TestGame_Start(t *testing.T) {

	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5, io.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7, io.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	t.Run("record ruth win ", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		dummyBlindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, store)
		winner := "Ruth"

		game.Finish(winner)
		poker.AssertPlayerWin(t, store, winner)
	})
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t testing.TB, blindAlerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		if len(blindAlerter.Alerts) <= i {
			t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
		}

		got := blindAlerter.Alerts[i]
		assertScheduledAlert(t, got, want)
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}
