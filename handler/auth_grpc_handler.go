package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

type AuthGrpcHandler struct {
	pb.UnimplementedAuthServer
	usecase usecase.UserUsecase
}

func NewAuthGrpcHandler(usecase usecase.UserUsecase) *AuthGrpcHandler {
	return &AuthGrpcHandler{
		usecase: usecase,
	}
}

func (h *AuthGrpcHandler) Login(ctx context.Context, data *pb.LoginReq) (*pb.LoginRes, error) {
	res, err := h.usecase.UserLogin(ctx, dto.LoginReq{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginRes{AccessToken: res.AccessToken}, nil
}
