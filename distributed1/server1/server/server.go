package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	userpb "server1/userpb"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedSendDataServer
}

func (s *server) Send(ctx context.Context, req *userpb.DataSend) (*userpb.Return, error) {
	make := fmt.Sprintf("files/%s", req.Save)
	t1 := req.Data
	file, err := os.Create(make)
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
	make := fmt.Sprintf("files/%s", req.FileName)
	chunkFile, err := os.Open(make)
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
	listener, err := net.Listen("tcp", "0.0.0.0:3001")
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
