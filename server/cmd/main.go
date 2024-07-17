package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	server "github.com/Jasspie/real-time-chat-app-v2/pkg"
	pb "github.com/Jasspie/real-time-chat-app-v2/proto/chat/v1"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}

	addr := args[0]
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(lis)

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	pb.RegisterChatServiceServer(s, &server.Server{
		RoomUsers: make(map[string]map[string]pb.ChatService_ChatServer),
		Rooms:     make(map[string]*pb.Room),
		Users:     make(map[string]*pb.User),
	})

	log.Printf("listening at %s\n", addr)

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
