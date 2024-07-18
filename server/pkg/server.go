package server

import (
	"context"
	"log"
	"sync"

	pb "github.com/Jasspie/real-time-chat-app-v2/proto/chat/v1"
)

type Server struct {
	RoomUsers map[string]map[string]pb.ChatService_ChatServer
	Rooms     map[string]bool
	Users     map[string]bool
	mu        sync.RWMutex
	pb.UnimplementedChatServiceServer
}

func (s *Server) addRoomUser(roomID string, userID string, srv pb.ChatService_ChatServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.RoomUsers[roomID][userID] = srv
}

func (s *Server) deleteRoomUser(roomID string, userID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.RoomUsers[roomID], userID)
	if len(s.RoomUsers[roomID]) == 0 {
		delete(s.RoomUsers, roomID)
	}
	delete(s.Rooms, roomID)
	delete(s.Users, userID)
}

func (s *Server) getRoomUserServers(roomID string) []pb.ChatService_ChatServer {
	var servers []pb.ChatService_ChatServer
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, server := range s.RoomUsers[roomID] {
		servers = append(servers, server)
	}
	return servers
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Users[req.Name]; ok {
		return &pb.CreateUserResponse{Error: 1, Msg: "user already exists"}, nil
	}
	s.Users[req.Name] = true
	return &pb.CreateUserResponse{Error: 0, Msg: "created new user"}, nil
}

func (s *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Rooms[req.Name]; ok {
		return &pb.CreateRoomResponse{Error: 0, Msg: "joined existing room"}, nil
	}
	s.Rooms[req.Name] = true
	return &pb.CreateRoomResponse{Error: 0, Msg: "created new room"}, nil
}

func (s *Server) GetRoomUsers(req *pb.GetRoomUsersRequest, stream pb.ChatService_GetRoomUsersServer) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for user := range s.RoomUsers[req.Room] {
		if err := stream.Send(&pb.GetRoomUsersResponse{User: user}); err != nil {
			log.Printf("could not get user %v", err)
		}
	}
	return nil
}

func (s *Server) Chat(req *pb.ChatRequest, stream pb.ChatService_ChatServer) error {
	if req.GetJoin() != nil { // process join request, add user to a room
		room := req.GetJoin().Room
		user := req.GetJoin().User
		s.addRoomUser(room, user, stream)
	} else if req.GetMsg() != nil { // process message, send message to all users in the room
		msg := req.GetMsg().Msg
		for _, server := range s.getRoomUserServers(msg.Room) {
			if err := server.Send(&pb.ChatResponse{Msg: msg}); err != nil {
				log.Printf("could not send message  %v", err)
			}
		}
	}
	return nil
}
