package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

type BookUsecase interface {
	GetAllBooks(query valueobject.Query) []*entity.Book
	GetSingleBook(query valueobject.Query) *entity.Book
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
func (u *bookUsecase) GetAllBooks(query valueobject.Query) []*entity.Book {
	return u.bookRepo.Find(query)
}

func (u *bookUsecase) GetSingleBook(query valueobject.Query) *entity.Book {
	return u.bookRepo.First(query)
}

func (u *bookUsecase) AddBook(book *entity.Book) (*entity.Book, error) {
	condition := *valueobject.NewCondition("title", valueobject.Equal, book.Title)
	query := valueobject.Query{
		Conditions: []valueobject.Condition{condition},
	}
	b := u.GetSingleBook(query)
	if b != nil {
		return nil, apperror.ErrAlreadyExist{
			Resource: "book",
			Field:    "title",
			Value:    b.Title,
		}
	}
	return u.bookRepo.Create(book)
}
