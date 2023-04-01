package db

import (
	"os"
	"testing"
)

var testStore *Store

func TestMain(m *testing.M) {

	testStore = NewStore()

	os.Exit(m.Run())
}
