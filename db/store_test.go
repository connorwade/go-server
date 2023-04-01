package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	store := testStore

	arg := CreateAlbumParams{
		Title:  "Ray",
		Artist: "l'arc-en-ciel",
		Price:  7.99,
	}

	album, err := store.Create(arg)

	assert.NoError(t, err)
	fmt.Println(album)
}

func TestAll(t *testing.T) {
	store := testStore

	fmt.Println(store)
}
