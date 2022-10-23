package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(f *os.File) (*FileSystemPlayerStore, error) {

	err := initialisePlayerDBFile(f)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(f)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", f.Name(), err)
	}
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&Tape{f}),
		league:   league,
	}, nil
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

func initialisePlayerDBFile(f *os.File) error {
	f.Seek(0, 0)
	info, err := f.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", f.Name(), err)
	}
	if info.Size() == 0 {
		f.Write([]byte("[]"))
		f.Seek(0, 0)
	}
	return nil
}
