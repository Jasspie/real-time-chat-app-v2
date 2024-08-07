package server

import (
	"context"
	"log"
	"sync"

	"connectrpc.com/connect"

	v1 "github.com/Jasspie/real-time-chat-app-v2/server/chat/v1"
	pb "github.com/Jasspie/real-time-chat-app-v2/server/chat/v1/v1connect"
)

type UserSession struct {
	Stream   *connect.ServerStream[v1.NewChatSessionResponse]
	UserName string
	RoomName string
	IsActive bool
	err      chan error
}

type ChatServer struct {
	pb.UnimplementedChatServiceHandler
	RoomUsers map[string][]*UserSession
	mu        sync.RWMutex
}

func (cr *ChatServer) NewChatSession(
	_ context.Context,
	req *connect.Request[v1.NewChatSessionRequest],
	stream *connect.ServerStream[v1.NewChatSessionResponse],
) error {
	session := &UserSession{
		Stream:   stream,
		UserName: req.Msg.UserName,
		RoomName: req.Msg.RoomName,
		IsActive: true,
		err:      make(chan error),
	}
	cr.mu.Lock()
	cr.RoomUsers[session.RoomName] = append(cr.RoomUsers[session.RoomName], session)
	cr.mu.Unlock()
	log.Printf("User %s joined room %s", session.UserName, session.RoomName)
	return <-session.err
}

func (cr *ChatServer) BroadcastChat(
	_ context.Context,
	req *connect.Request[v1.BroadcastChatRequest],
) (
	*connect.Response[v1.BroadcastChatResponse],
	error,
) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	wait := sync.WaitGroup{}
	done := make(chan int)
	for _, conn := range cr.RoomUsers[req.Msg.Msg.RoomName] {
		wait.Add(1)

		go func(msg *v1.Msg, conn *UserSession) {
			defer wait.Done()

			if conn.IsActive {
				err := conn.Stream.Send(&v1.NewChatSessionResponse{Msg: msg})
				log.Printf("Sending message to: %v from %v\n", conn.UserName, msg.UserName)

				if err != nil {
					log.Printf("Error sending message: %v\n", err)
					conn.IsActive = false
					conn.err <- err
				}
			}
		}(req.Msg.Msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &connect.Response[v1.BroadcastChatResponse]{}, nil
}
