package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	userpb "server3/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedSendDataServer
}

func (s *server) Send(ctx context.Context, req *userpb.DataSend) (*userpb.Return, error) {
	t1 := req.Data
	file, err := os.Create(req.Save)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(t1) // Convert byte to a slice of one element
	if err != nil {
		panic(err)
	}
	return &userpb.Return{
		Status: true,
	}, nil
}
func (s *server) Get(ctx context.Context, req *userpb.GetDataSend) (*userpb.GetDataReturn, error) {
	chunkFile, err := os.Open(req.FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer chunkFile.Close()

	// Read chunk file into byte slice
	chunkData, err := io.ReadAll(chunkFile)
	if err != nil {
		panic(err)
	}
	return &userpb.GetDataReturn{
		Status: true,
		Data:   chunkData,
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterSendDataServer(grpcServer, &server{})

	fmt.Println("gRPC Server is running on port 3001...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
