package cmd

import (
	"context"
	"fmt"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func StartGrpcClient() {
	// Set up the gRPC client to connect to the server
	conn, err := grpc.Dial("10.20.191.148:50051", grpc.WithInsecure())
	// Tazki 163
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewBorrowClient(conn)

	// Prepare the request message
	req := &pb.BorrowRequest{
		BookId: 1,
		UserId: 6,
	}

	ctx := context.Background()
	md := metadata.Pairs("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJpc3MiOiJhdXRoVG9rZW4iLCJleHAiOjE3MDIzNzc2MDgsImlhdCI6MTcwMjM3NDAwOH0.CRWgtEw2dnmKKbtZ5o2tNVbma1z_sb4uStuwESaXB1o")
	// md := metadata.Pairs("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDI0NjA2MDgsImlhdCI6MTcwMjM3NDIwOCwiaXNzIjoiTElCUkFSWS1BUEkiLCJ1c2VyX2lkIjoyfQ.iaxo9ohH-_L1mVdcQHuylyIFu-QSiRGHLlIAjMWz044")
	ctx2 := metadata.NewOutgoingContext(ctx, md)

	// Send the request to the server
	res, err := client.Borrow(ctx2, req)
	if err != nil {
		log.Fatalf("Failed to send the request: %v", err)
	}

	// Print the server's response
	fmt.Println("Server Response: ", res)
}
