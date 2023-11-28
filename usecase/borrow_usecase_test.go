package usecase_test

// import (
// 	"testing"

// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/dto"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/mocks"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
// 	"github.com/go-playground/assert/v2"
// )

// var borrows = []models.BorrowBook{
// 	{
// 		UserId: 1,
// 		BookId: 1,
// 		Status: "not returned",
// 	},
// }

// var borrowsReq = []dto.BorrowReq{
// 	{
// 		UserId: 1,
// 		BookId: 1,
// 	},
// }

// var borrowsRes = []dto.BorrowRes{
// 	{
// 		UserId: 1,
// 		BookId: 1,
// 		Status: "not returned",
// 	},
// }

// func TestGetAllRecords(t *testing.T) {
// 	t.Run("should return borrow records when success", func(t *testing.T) {
// 		borrowRepository := mocks.NewBorrowRepository(t)
// 		bookRepository := mocks.NewBookRepository(t)
// 		userRepository := mocks.NewUserRepository(t)
// 		borrowUsecase := usecase.NewBorrowUsecase(borrowRepository, bookRepository, userRepository)
// 		borrowRepository.On("FindBorrows").Return(borrows, nil)

// 		resBorrows, _ := borrowUsecase.GetAllRecords()

// 		assert.Equal(t, borrows, resBorrows)
// 	})
// }

// func TestBorrowBook(t *testing.T) {
// t.Run("should return newborrow record when success", func(t *testing.T) {
// 	borrowRepository := mocks.NewBorrowRepository(t)
// 	bookRepository := mocks.NewBookRepository(t)
// 	userRepository := mocks.NewUserRepository(t)
// 	borrowUsecase := usecase.NewBorrowUsecase(borrowRepository, bookRepository, userRepository)
// 	userRepository.On("FindUserById", mock.Anything).Return(nil, errors.New("Fake error"))

// 	resBorrows, _ := borrowUsecase.BorrowBook(borrows[0])

// 	assert.Equal(t, borrows, resBorrows)
// })
// }
