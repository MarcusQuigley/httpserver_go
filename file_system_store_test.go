package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo","Wins":10},
			{"Name": "Chris","Wins":33}]`)

		store := FileSystemPlayerStore{database}
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cletus","Wins":10},
			{"Name": "Klaus","Wins":20}
		]`)
		player := "Klaus"
		store := FileSystemPlayerStore{database}

		got, err := store.GetPlayerScore(player)
		if err != nil {
			t.Fatalf("got error %v", err)
		}

		want := 20
		assertScoreEquals(t, got, want)

	})
}
