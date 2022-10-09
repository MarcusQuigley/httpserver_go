package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	f.database.Seek(0, 0)
	league, err := NewLeague(f.database)
	if err != nil {
		return -1, err
	}
	var wins int
	for _, v := range league {
		if v.Name == name {
			wins = v.Wins
			break
		}
	}
	return wins, nil // errors.New("no player found with " + name)
}
