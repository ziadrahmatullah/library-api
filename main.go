package main

import (
	"log"
	"net/http"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	db := database.ConnectDB()

	addr := os.Getenv("APP_PORT")
	br := repository.NewBookRepository(db)
	bu := usecase.NewBookUsecase(br)
	bh := handler.NewBookHandler(bu)

	opts := server.RouterOpts{
		ProductHandler: bh,
	}
	r := server.NewRouter(opts)

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	srv.ListenAndServe()
}
