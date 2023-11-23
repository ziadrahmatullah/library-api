package main

import (
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
)

func main() {
	db, err := repository.GetConnection()
	if err != nil {
		log.Println(err)
	}
	b := &entity.Book{}
	_ = db.AutoMigrate(b)

	books := []*entity.Book{
		{Id: 1, Title: "How to eat", Description: "Explain how to eat", Quantity: 2, Cover: "kertas"},
	}
	db.Create(books)
}
