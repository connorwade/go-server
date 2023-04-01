package db

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"
)

type Store struct {
	db []*Music `json:"music"`
}

func NewStore() *Store {
	return &Store{
		db: []*Music{
			{ID: uuid.New(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: uuid.New(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (s *Store) Migrate() (*[]Music, error) {
	var migratedDB []Music

	if _, err := os.Stat("db.json"); err != nil {

		migratedDB := []Music{
			{ID: uuid.New(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: uuid.New(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		}

		file, err := json.MarshalIndent(migratedDB, "", " ")
		if err != nil {
			return nil, err
		}

		err = os.WriteFile("db.json", file, 0644)
		if err != nil {
			return nil, err
		}
	}

	content, err := os.ReadFile("db.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, []Music{})

	return &migratedDB, err
}

type CreateAlbumParams struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (s *Store) Create(arg CreateAlbumParams) (*Music, error) {

	for _, v := range s.db {
		if v.Title == arg.Title && v.Artist == v.Artist {
			return nil, errors.New("this item already exists")
		}
	}

	newMusic := &Music{
		ID:     uuid.New(),
		Title:  arg.Title,
		Artist: arg.Artist,
		Price:  arg.Price,
	}

	s.db = append(s.db, newMusic)

	file, err := json.MarshalIndent(s.db, "", " ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("db.json", file, 0644)
	if err != nil {
		return nil, err
	}

	return newMusic, nil
}

func (s *Store) All() ([]*Music, error) {
	// content, err := os.ReadFile("db.json")

	// if err != nil {
	// 	return nil, err
	// }

	// err = json.Unmarshal(content, &s.db)

	return s.db, nil
}

type GetAlbumByIDArgs struct {
	ID uuid.UUID `json:"id"`
}

func (s *Store) GetByID(arg GetAlbumByIDArgs) (*Music, error) {
	for _, v := range s.db {
		if v.ID == arg.ID {
			return v, nil
		}
	}
	return nil, errors.New("cannot find album with that id")
}

type UpdateAlbumArguments struct {
	id       uuid.UUID `json:"id"`
	newMusic *Music    `json:"newMusic"`
}

func (s *Store) Update(arg UpdateAlbumArguments) (*Music, error) {
	for i, v := range s.db {
		if v.ID == arg.id {
			s.db[i] = arg.newMusic
			file, err := json.MarshalIndent(s.db, "", " ")
			if err != nil {
				return nil, err
			}

			err = os.WriteFile("db.json", file, 0644)
			if err != nil {
				return nil, err
			}

			return arg.newMusic, nil
		}
	}
	return nil, errors.New("cannot find album with that id")
}

type DeleteAlbumArguments struct {
	id uuid.UUID `json:"id"`
}

func (s *Store) Delete(arg DeleteAlbumArguments) error {
	for i, v := range s.db {
		if v.ID == arg.id {
			s.db = append(s.db[:i], s.db[i+1:]...)
			file, err := json.MarshalIndent(s.db, "", " ")
			if err != nil {
				return err
			}

			err = os.WriteFile("db.json", file, 0644)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("cannot find album with that id")
}
