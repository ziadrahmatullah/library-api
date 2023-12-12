package cmd

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/middleware"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/usecase"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pb "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
)

func StartGrpcServer() {
	db := database.ConnectDB()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	authGrpcHandler := handler.NewAuthGrpcHandler(uu)
	userGrpcHandler := handler.NewUserGrpcHandler(uu)

	br := repository.NewBookRepository(db)
	bu := usecase.NewBookUsecase(br)
	bookGrpcHandler := handler.NewBookGrpcHandler(bu)

	bbr := repository.NewBorrowRepository(db)
	bbu := usecase.NewBorrowUsecase(bbr, br, ur)
	borrowGrpcHandler := handler.NewBorrowGrpcHandler(bbu)

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Err(err).Msg("error starting tcp server")
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		middleware.LoggerInterceptor,
		middleware.ErrorInterceptor,
		middleware.AuthInterceptor,
	))

	pb.RegisterAuthServer(server, authGrpcHandler)
	pb.RegisterUserHandlerServer(server, userGrpcHandler)
	pb.RegisterBookServiceServer(server, bookGrpcHandler)
	pb.RegisterBorrowServer(server, borrowGrpcHandler)

	log.Info().Msg("starting grpc server")

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(list); err != nil {
			signCh <- syscall.SIGINT
		}
	}()
	log.Info().Msg("server started")
	<-signCh
	log.Info().Msg("server stopped")

}
