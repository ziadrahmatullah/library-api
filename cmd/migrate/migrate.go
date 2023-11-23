package main

import (
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/migration"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
)

func main() {
	db, err := repository.GetConnection()
	if err != nil {
		log.Println(err)
	}
	migration.Migrate(db)
}
