package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type BorrowUsecase interface {
	BorrowBook(*models.BorrowingBook) (*models.BorrowingBook, error)
}

type borrowUsecase struct {
	borrowRepository repository.BorrowRepository
}

func NewBorrowUsecase(bu repository.BorrowRepository) BorrowUsecase {
	return &borrowUsecase{
		borrowRepository: bu,
	}
}

func (bu *borrowUsecase) BorrowBook(borrow *models.BorrowingBook) (*models.BorrowingBook, error) {
	return bu.borrowRepository.BorrowBook(borrow)
}
