package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(f *os.File) *FileSystemPlayerStore {
	f.Seek(0, 0)
	league, _ := NewLeague(f)
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&Tape{f}),
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	//league := f.GetLeague()
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)

}
