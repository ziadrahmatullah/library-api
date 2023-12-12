package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/appvalidator"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

type AuthGrpcHandler struct {
	pb.UnimplementedAuthServer
	usecase   usecase.UserUsecase
	validator appvalidator.AppValidator
}

func NewAuthGrpcHandler(usecase usecase.UserUsecase, val appvalidator.AppValidator) *AuthGrpcHandler {
	return &AuthGrpcHandler{
		usecase: usecase,
		validator: val,
	}
}

func (h *AuthGrpcHandler) Login(ctx context.Context, data *pb.LoginReq) (*pb.LoginRes, error) {
	userReq := dto.LoginReq{
		Email:    data.Email,
		Password: data.Password,
	}
	err := h.validator.Validate(userReq)
	if err != nil {
		return nil, apperror.ErrInvalidBody
	}

	res, err := h.usecase.UserLogin(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return &pb.LoginRes{AccessToken: res.AccessToken}, nil
}
