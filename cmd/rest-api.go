package cmd

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

func StartRestAPI() {
	db := database.ConnectDB()

	addr := os.Getenv("APP_PORT")
	br := repository.NewBookRepository(db)
	bu := usecase.NewBookUsecase(br)
	bh := handler.NewBookHandler(bu)

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	brr := repository.NewBorrowRepository(db)
	bru := usecase.NewBorrowUsecase(brr, br, ur)
	brh := handler.NewBorrowHandler(bru)

	opts := server.RouterOpts{
		BookHandler:   bh,
		UserHandler:   uh,
		BorrowHandler: brh,
	}

	// db.Debug().AutoMigrate(&models.Book{}, &models.Author{}, &models.BorrowBook{}, &models.User{})
	r := server.NewRouter(opts)

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	// Wait forloggerwithconfig interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
