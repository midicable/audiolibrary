package main

import (
	"audiolibrary/client/pb"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	title = flag.String("title", "none", "book title")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAudiolibraryClient(conn)
	filepath, err := client.GetAudiofilePath(context.Background(), &pb.AudiolibraryRequest{BookTitle: *title})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println(filepath)
}
