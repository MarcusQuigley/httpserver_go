package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

// FileSystemPlayerStoreFromFile creates a PlayerStore from the contents of a JSON file found at path.
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s, %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return store, closeFunc, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
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
