package server

// import (
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
// 	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
// 	"gorm.io/gorm"
// )

// func Handlers(db *gorm.DB) (opts *RouterOpts) {
// 	opts = &RouterOpts{
// 		registerBookHandler(db),
// 		registerUserHandler(db),
// 		registerBorrowBookHandler(db),
// 	}
// 	return
// }

// func registerBookHandler(db *gorm.DB) (bookHandler *handler.BookHandler) {
// 	bookRepository := repository.NewBookRepository(db)
// 	bookService := usecase.NewBookUsecase(bookRepository)
// 	bookHandler = handler.NewBookHandler(bookService)
// 	return
// }

// func registerUserHandler(db *gorm.DB) (userHandler *handler.UserHandler) {
// 	userRepository := repository.NewUserRepository(db)
// 	userService := usecase.NewUserUsecase(userRepository)
// 	userHandler = handler.NewUserHandler(userService)
// 	return
// }

// func registerBorrowBookHandler(db *gorm.DB) (borrowBookHandler *handler.BorrowBookHandler) {
// 	borrowBookRepository := repository.NewBorrowRepository(db)
// 	borrowBookService := usecase.NewBorrowUsecase(borrowBookRepository)
// 	borrowBookHandler = handler.NewBorrowBookHandler(borrowBookService)
// 	return
// }