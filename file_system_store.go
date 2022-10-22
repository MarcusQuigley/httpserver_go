package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Writer //.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(d io.ReadWriteSeeker) *FileSystemPlayerStore {
	d.Seek(0, 0)
	league, _ := NewLeague(d)
	return &FileSystemPlayerStore{&Tape{d}, league}
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
	json.NewEncoder(f.database).Encode(f.league)

}
