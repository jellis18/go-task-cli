package main

import (
	"log"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"
)

func main() {
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
