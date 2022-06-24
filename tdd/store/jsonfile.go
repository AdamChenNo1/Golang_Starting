/*
 * File: /store/jsonfile.go                                                    *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 11:12:37                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 13:26:58                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package store

import (
	"encoding/json"
	"fmt"
	"go_start/tdd/model"
	"io"
	"os"
)

type FileSystemStore struct {
	database *json.Encoder
	league   model.League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemStore, error) {
	err := initialisePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialsing player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store form file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{json.NewEncoder(&tape{file}), league}, nil
}

func (f *FileSystemStore) GetLeague() model.League {
	return f.league
}

func NewLeague(rdr io.Reader) (model.League, error) {
	var league model.League

	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, model.Player{name, 1})
	}

	f.database.Encode(&f.league)
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}
