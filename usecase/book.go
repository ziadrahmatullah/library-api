package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BookUsecase interface {
	GetAllBooks(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book
	GetSingleBook(conditions []valueobject.Condition) *entity.Book
	AddBook(book *entity.Book) (*entity.Book, error)
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: repo,
	}
}
func (u *bookUsecase) GetAllBooks(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book {
	return u.bookRepo.Find(clause, conditions)
}

func (u *bookUsecase) GetSingleBook(conditions []valueobject.Condition) *entity.Book {
	return u.bookRepo.First(conditions)
}

func (u *bookUsecase) AddBook(book *entity.Book) (*entity.Book, error) {
	condition := *valueobject.NewCondition("title", valueobject.Equal, book.Title)
	b := u.GetSingleBook([]valueobject.Condition{condition})
	if book != nil {
		return nil, apperror.ErrAlreadyExist{
			Resource: "book",
			Field:    "title",
			Value:    b.Title,
		}
	}
	return u.bookRepo.Create(book)
}
