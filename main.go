package main

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

func main() {
	db := database.ConnectDB()

	// addr := os.Getenv("APP_PORT")
	br := repository.NewBookRepository(db)
	bu := usecase.NewBookUsecase(br)
	bh := handler.NewBookHandler(bu)

	opts := server.RouterOpts{
		ProductHandler: bh,
	}
	r := server.NewRouter(opts)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	srv.ListenAndServe()
}
