package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/appjwt"
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

	jwt := appjwt.NewJwt()
	authUsecase := usecase.NewAuthUsecase(userRepo, jwt)
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

	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	const defaultTimeout = 5
	timeoutString := os.Getenv("TIMEOUT")
	timeout, err := strconv.Atoi(timeoutString)
	if err != nil {
		timeout = defaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
