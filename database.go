package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type database struct {
	Data Data
}

type Repository interface {
	Migrate() error
	Create(music Music) error
	All() ([]Music, error)
	GetByID(id string) (*Music, error)
	Update(id string, newMusic Music) (*Music, error)
	Delete(id string) error
}

func (db *database) Migrate() error {
	if _, err := os.Stat("db.json"); err != nil {
		migratedDB := Data{
			Albums: []Music{
				{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
				{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
				{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
			},
		}

		file, err := json.MarshalIndent(migratedDB, "", " ")
		if err != nil {
			return err
		}

		err = ioutil.WriteFile("db.json", file, 0644)
		if err != nil {
			return err
		}
	}

	content, err := ioutil.ReadFile("db.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &db.Data)

	return err
}

func (db *database) Create(music Music) error {

	for _, v := range db.Data.Albums {
		if v.ID == music.ID {
			return errors.New("this item already exists")
		}
	}

	db.Data.Albums = append(db.Data.Albums, music)

	file, err := json.MarshalIndent(db.Data, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("db.json", file, 0644)
	return err
}

func (db *database) All() ([]Music, error) {
	content, err := ioutil.ReadFile("db.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &db.Data)

	return db.Data.Albums, err
}

func (db *database) GetByID(id string) (*Music, error) {
	for _, v := range db.Data.Albums {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errors.New("cannot find album with that id")
}

func (db *database) Update(id string, newMusic Music) (*Music, error) {
	for i, v := range db.Data.Albums {
		if v.ID == id {
			db.Data.Albums[i] = newMusic
			file, err := json.MarshalIndent(db.Data, "", " ")
			if err != nil {
				return nil, err
			}

			err = ioutil.WriteFile("db.json", file, 0644)
			if err != nil {
				return nil, err
			}

			return &newMusic, nil
		}
	}
	return nil, errors.New("cannot find album with that id")
}

func (db *database) Delete(id string) error {
	for i, v := range db.Data.Albums {
		if v.ID == id {
			db.Data.Albums = append(db.Data.Albums[:i], db.Data.Albums[i+1:]...)
			file, err := json.MarshalIndent(db.Data, "", " ")
			if err != nil {
				return err
			}

			err = ioutil.WriteFile("db.json", file, 0644)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("cannot find album with that id")
}
