package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
)

type BookGrpcHandler struct {
	pb.UnimplementedBookServiceServer
	bu usecase.BookUsecase
}

func NewBookGrpcHandler(bu usecase.BookUsecase) *BookGrpcHandler {
	return &BookGrpcHandler{
		bu: bu,
	}
}

func (h *BookGrpcHandler) GetAllBook(ctx context.Context, req *pb.BooksReq) (*pb.BooksRes, error) {
	res, err := h.bu.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	var bookRes pb.BooksRes
	for _, book := range res {
		var bookGrpc pb.Book
		bookGrpc.Id = uint32(book.ID)
		bookGrpc.CreatedAt = book.CreatedAt.String()
		bookGrpc.UpdatedAt = book.UpdatedAt.String()
		bookGrpc.DeletedAt = book.DeletedAt.Time.String()
		bookGrpc.Title = book.Title
		bookGrpc.Description = book.Description
		bookGrpc.Quantity = uint32(book.Quantity)
		bookGrpc.Cover = book.Cover
		bookGrpc.AuthorId = uint32(book.AuthorId)
		bookRes.Books = append(bookRes.Books, &bookGrpc)
	}
	return &bookRes, nil
}

func (h *BookGrpcHandler) CreateBook(ctx context.Context, req *pb.CreateBookReq) (*pb.Book, error) {
	qty := int(req.Quantity)
	bookReq := dto.BookReq{
		Title:       req.Title,
		Description: req.Description,
		Quantity:    &qty,
		Cover:       req.Cover,
		AuthorId:    uint(req.AuthorId),
	}
	res, err := h.bu.CreateBook(ctx, bookReq.ToBookModel())
	if err != nil {
		return nil, err
	}
	return &pb.Book{
		Id:          uint32(res.ID),
		CreatedAt:   res.CreatedAt.String(),
		UpdatedAt:   res.UpdatedAt.String(),
		DeletedAt:   res.DeletedAt.Time.String(),
		Title:       res.Title,
		Description: res.Description,
		Quantity:    uint32(res.Quantity),
		Cover:       res.Cover,
		AuthorId:    uint32(res.AuthorId),
	}, nil
}
