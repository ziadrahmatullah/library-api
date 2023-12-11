package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb/proto"

	"google.golang.org/protobuf/proto"
)

func StartGrpcServer() {
	user := &pb.User{
		Id:       1,
		Name:     "alice",
		Email:    "alice@gmail.com",
		Phone:    "081234567890",
		Password: "alice123",
	}

	userJson := &models.User{
		Name:     "alice",
		Email:    "alice@gmail.com",
		Phone:    "081234567890",
		Password: "alice123",
	}

	data, _ := proto.Marshal(user)
	err := os.WriteFile("./response/protobuff", data, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	dataJSON, _ := json.Marshal(userJson)
	err = os.WriteFile("./response/response.json", dataJSON, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
