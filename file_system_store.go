package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
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

func (f *FileSystemPlayerStore) RecordWin(name string) {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)

	for i, v := range league {
		if v.Name == name {
			league[i].Wins += 1
			break
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)

}
