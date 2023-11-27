package main

import (
	"log"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/router"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
)

func main() {
	db, err := repository.GetConnection()
	if err != nil {
		log.Println(err)
	}

	authorRepo := repository.NewAuthorRepository(db)

	bookRepo := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo, authorRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	borrowingRecordRepo := repository.NewBorrowingRecordsRepository(db)
	borrowingRecordUsecase := usecase.NewBorrowingRecordUsecase(borrowingRecordRepo, bookRepo)
	borrowingRecordHandler := handler.NewBorrowingRecordHandler(borrowingRecordUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	handlers := router.Handlers{
		Book:            bookHandler,
		User:            userHandler,
		BorrowingRecord: borrowingRecordHandler,
		Auth:            authHandler,
	}
	r := router.New(handlers)

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
