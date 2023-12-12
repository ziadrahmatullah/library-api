package main

import (
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	// cmd.StartRestAPI()
	cmd.StartGrpcServer()
	// cmd.StartGrpcClient()
}
