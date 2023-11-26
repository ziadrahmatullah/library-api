package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
)

type BorrowUsecase interface {
	GetAllRecords()([]models.BorrowBook, error)
	BorrowBook(models.BorrowBook) (*models.BorrowBook, error)
	ReturnBook(models.BorrowBook) (*models.BorrowBook, error)
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

func (bu *borrowUsecase) GetAllRecords() ([]models.BorrowBook, error){
	return bu.borrowRepository.FindBorrows()
}

func (bu *borrowUsecase) BorrowBook(borrow models.BorrowBook) (*models.BorrowBook, error) {
	if user, err := bu.userRepository.FindUserById(borrow.UserId); user == nil{
		return nil, err
	}
	book, err := bu.bookRepository.FindBooksById(borrow.BookId)
	if book == nil {
		return nil, err
	}else if book.Quantity == 0{
		return nil, apperror.ErrBookOutOfStock
	}
	if err := bu.bookRepository.DecreaseBookQty(borrow.BookId); err != nil{
		return nil, err
	}
	return bu.borrowRepository.NewBorrow(borrow)
}

func (bu *borrowUsecase) ReturnBook(borrow models.BorrowBook) (*models.BorrowBook, error){
	id, err := bu.borrowRepository.FindBorrow(borrow)
	if id == 0{
		return nil, err
	}
	if err := bu.bookRepository.IncreaseBookQty(borrow.BookId); err != nil{
		return nil, err
	}
	return bu.borrowRepository.UpdateBorrowStatus(id)	
}
