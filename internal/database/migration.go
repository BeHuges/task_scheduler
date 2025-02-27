package database

import (
	"log"
	"os"
	"path/filepath"

	"task_scheduler/internal/repository"
)

func Migration(rep *repository.Repository) {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), os.Getenv("TODO_DBFILE"))
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	if install {
		if err := rep.CreateTable(); err != nil {
			log.Fatal(err)
		}
	}
}
