package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

type BorrowGrpcHandler struct {
	pb.UnimplementedBorrowServer
	bu usecase.BorrowUsecase
}

func NewBorrowGrpcHandler(bu usecase.BorrowUsecase) *BorrowGrpcHandler {
	return &BorrowGrpcHandler{
		bu: bu,
	}
}

func (h *BorrowGrpcHandler) Borrow(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowResponse, error) {
	borrowReq := dto.BorrowReq{
		BookId: uint(req.BookId),
	}

	// if ctx.Value("id") != req.UserId {
	// 	log.Println("here")
	// 	return &pb.BorrowResponse{}, apperror.ErrBookNotFound
	// }
	res, err := h.bu.BorrowBook(ctx, borrowReq.ToBorrowModel(uint(req.UserId)))
	if err != nil {
		return nil, err
	}
	return &pb.BorrowResponse{
		Id:     uint32(res.ID),
		UserId: uint32(res.UserId),
		BookId: uint32(res.BookId),
		Status: res.Status,
	}, nil
}
