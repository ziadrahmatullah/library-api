package main

import (
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"gorm.io/gorm"
)

func main() {
	db, err := repository.GetConnection()
	if err != nil {
		log.Println(err)
	}
	migrate(db)
	seed(db)
}

func migrate(db *gorm.DB) {
	b := &entity.Book{}

	_ = db.Migrator().DropTable(b)

	_ = db.AutoMigrate(b)
}

func seed(db *gorm.DB) {
	books := []*entity.Book{
		{Title: "How to eat", Description: "Explain how to eat", Quantity: 2, Cover: "kertas"},
	}
	db.Create(books)
}
