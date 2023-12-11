package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

type UserGrpcHandler struct{
	pb.UnimplementedUserHandlerServer
	uu usecase.UserUsecase
}

func NewUserGrpcHandler(uu usecase.UserUsecase) *UserGrpcHandler{
	return &UserGrpcHandler{
		uu : uu,
	}
}

func (h *UserGrpcHandler) GetAllUsers(ctx context.Context, req *pb.UsersReq) (*pb.UsersRes, error){
	res, err := h.uu.GetAllUsers(ctx)
	if err != nil{
		return nil, err
	}
	var userRes pb.UsersRes
	for _, user := range res{
		var userGrpc pb.User
		userGrpc.Id = uint32(user.ID)
		userGrpc.Name = user.Name
		userGrpc.Email = user.Email
		userGrpc.Phone = user.Phone
		userGrpc.Password = user.Password
		userRes.Users = append(userRes.Users, &userGrpc)
	}
	return &userRes, nil
}