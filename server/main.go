package main

import (
	"audiolibrary/server/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedAudiolibraryServer
}

type Audiobook struct {
	ID    int
	Title string
	Path  string
}

var database *sql.DB

func (s server) GetAudiofilePath(ctx context.Context, in *pb.AudiolibraryRequest) (*pb.AudiolibraryResponse, error) {
	log.Printf("Received book title: %v", in.GetBookTitle())
	if err != nil {
		log.Println(err)
	}
	
	return &pb.AudiolibraryResponse {
		FilePath: "somefilepath to " + in.GetBookTitle(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Printf("starting server on localhost:8080...")
	s := grpc.NewServer()
	pb.RegisterAudiolibraryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
