package server

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/google/uuid"

	pb "github.com/Jasspie/real-time-chat-app-v2/proto/chat/v1"
)

type Server struct {
	RoomUsers map[string]map[string]pb.ChatService_ChatServer
	Rooms     map[string]*pb.Room
	Users     map[string]*pb.User
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
		return &pb.CreateUserResponse{Error: 0, Msg: "user already exists"}, nil
	}
	userID, err := uuid.NewRandom()
	if err != nil {
		return &pb.CreateUserResponse{Error: 1, Msg: "could not create new user ID"}, nil
	}
	s.Users[userID.String()] = &pb.User{Id: userID.String(), Name: req.Name}
	return &pb.CreateUserResponse{Error: 0, Msg: "created new user"}, nil
}

func (s *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Rooms[req.Name]; ok {
		return &pb.CreateRoomResponse{Error: 0, Msg: "joined existing room"}, nil
	}
	roomID, err := uuid.NewRandom()
	if err != nil {
		return &pb.CreateRoomResponse{Error: 1, Msg: "could not create new room ID"}, nil
	}
	s.Rooms[roomID.String()] = &pb.Room{Id: roomID.String(), Name: req.Name}
	return &pb.CreateRoomResponse{Error: 0, Msg: "created new room"}, nil
}

func (s *Server) GetRoomUsers(req *pb.GetRoomUsersRequest, stream pb.ChatService_GetRoomUsersServer) error {
	roomID := req.RoomId
	s.mu.RLock()
	defer s.mu.RUnlock()
	for userID := range s.RoomUsers[roomID] {
		if err := stream.Send(&pb.GetRoomUsersResponse{User: s.Users[userID]}); err != nil {
			log.Printf("could not get  %v", err)
			return err
		}
	}
	return nil
}

func (s *Server) Chat(stream pb.ChatService_ChatServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if req.GetMsg() != nil { // process message, send message to all Users in the room
			msg := req.GetMsg().Msg
			roomID := msg.RoomId
			for _, server := range s.getRoomUserServers(roomID) {
				if err := server.Send(&pb.ChatResponse{Msg: msg}); err != nil {
					log.Printf("broadcast err: %v", err)
				}
			}
		} else if req.GetJoin() != nil { // process join request, add user to a room
			roomID := req.GetJoin().RoomId
			userID := req.GetJoin().UserId
			s.addRoomUser(roomID, userID, stream)
			defer s.deleteRoomUser(roomID, userID)
		}
	}
}
