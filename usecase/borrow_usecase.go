package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type BorrowUsecase interface {
	GetAllRecords(context.Context)([]models.BorrowBook, error)
	BorrowBook(context.Context, models.BorrowBook) (*models.BorrowBook, error)
	ReturnBook(context.Context, models.BorrowBook) (*models.BorrowBook, error)
}

type borrowUsecase struct {
	borrowRepository repository.BorrowRepository
	bookRepository repository.BookRepository
	userRepository repository.UserRepository
}

func NewBorrowUsecase(brr repository.BorrowRepository, br repository.BookRepository, ur repository.UserRepository) BorrowUsecase {
	return &borrowUsecase{
		borrowRepository: brr,
		bookRepository: br,
		userRepository: ur,
	}
}

func (bu *borrowUsecase) GetAllRecords(ctx context.Context) ([]models.BorrowBook, error){
	return bu.borrowRepository.FindBorrows(ctx)
}

func (bu *borrowUsecase) BorrowBook(ctx context.Context, borrow models.BorrowBook) (*models.BorrowBook, error) {
	if user, err := bu.userRepository.FindUserById(ctx, borrow.UserId); user == nil{
		return nil, err
	}
	book, err := bu.bookRepository.FindBooksById(ctx, borrow.BookId)
	if book == nil {
		return nil, err
	}else if book.Quantity == 0{
		return nil, apperror.ErrBookOutOfStock
	}
	return bu.borrowRepository.NewBorrow(ctx, borrow)
}

func (bu *borrowUsecase) ReturnBook(ctx context.Context, borrow models.BorrowBook) (*models.BorrowBook, error){
	id, err := bu.borrowRepository.FindBorrow(ctx, borrow)
	if id == 0{
		return nil, err
	}
	return bu.borrowRepository.UpdateBorrowStatus(ctx, id)	
}
