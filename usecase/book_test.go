package usecase_test

import (
	"context"
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var author = &entity.Author{
	Id:   1,
	Name: "Alice",
}
var book = &entity.Book{
	Id:          1,
	Title:       "Title",
	Quantity:    1,
	Description: "Description",
	Cover:       "cover",
}
var books = []*entity.Book{book}

type BookUsecaseTestSuite struct {
	suite.Suite
	bookRepo    *mocks.BookRepository
	authorRepo  *mocks.AuthorRepository
	bookUsecase usecase.BookUsecase
}

func (s *BookUsecaseTestSuite) SetupSubTest() {
	s.bookRepo = mocks.NewBookRepository(s.T())
	s.authorRepo = mocks.NewAuthorRepository(s.T())
	s.bookUsecase = usecase.NewBookUsecase(s.bookRepo, s.authorRepo)
}

func (s *BookUsecaseTestSuite) TestBookUsecase_GetAllBooks() {
	s.Run("should return books", func() {
		query := valueobject.NewQuery()
		s.bookRepo.On("Find", mock.Anything, mock.Anything).Return(books, nil)

		fetchedBooks, err := s.bookUsecase.GetAllBooks(context.Background(), query)

		s.Equal(books, fetchedBooks)
		s.NoError(err)
	})
}

func (s *BookUsecaseTestSuite) TestBookUsecase_GetSingeBook() {
	s.Run("should return book", func() {
		query := valueobject.NewQuery()
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(book, nil)

		fetchedBook, err := s.bookUsecase.GetSingleBook(context.Background(), query)

		s.Equal(book, fetchedBook)
		s.NoError(err)
	})
}

func (s *BookUsecaseTestSuite) TestBookUsecase_AddBook() {
	s.Run("should return newly created book", func() {
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.authorRepo.On("First", mock.Anything, mock.Anything).Return(author, nil)
		s.bookRepo.On("Create", mock.Anything, mock.Anything).Return(book, nil)

		fetchedBook, err := s.bookUsecase.AddBook(context.Background(), book)

		s.Equal(book, fetchedBook)
		s.NoError(err)
	})
	s.Run("should return error when there's an error when searching book", func() {
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		fetchedBook, err := s.bookUsecase.AddBook(context.Background(), book)

		s.Nil(fetchedBook)
		s.Error(err)
	})
	s.Run("should return error when there's a book with same title", func() {
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(book, nil)

		fetchedBook, err := s.bookUsecase.AddBook(context.Background(), book)

		s.Nil(fetchedBook)
		s.Error(err)
	})
	s.Run("should return error when there's an error when searching author", func() {
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.authorRepo.On("First", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		fetchedBook, err := s.bookUsecase.AddBook(context.Background(), book)

		s.Nil(fetchedBook)
		s.Error(err)
	})
	s.Run("should return error when the author doesn't exists", func() {
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.authorRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)

		fetchedBook, err := s.bookUsecase.AddBook(context.Background(), book)

		s.Nil(fetchedBook)
		s.Error(err)
	})
	s.Run("should return error when there's an error when creating new book", func() {
		s.bookRepo.On("First", mock.Anything, mock.Anything).Return(nil, nil)
		s.authorRepo.On("First", mock.Anything, mock.Anything).Return(author, nil)
		s.bookRepo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		fetchedBook, err := s.bookUsecase.AddBook(context.Background(), book)

		s.Nil(fetchedBook)
		s.Error(err)
	})
}

func TestBookUsecase(t *testing.T) {
	suite.Run(t, new(BookUsecaseTestSuite))
}
