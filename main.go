package main

import (
	"github.com/connorwade/go-server/db"
	"github.com/connorwade/go-server/server"
)

func main() {
	store := db.NewStore()
	s := server.NewServer(store)
	s.Start("localhost:8080")
}
