package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	chatv1connect "github.com/Jasspie/real-time-chat-app-v2/chat/v1/v1connect"
	server "github.com/Jasspie/real-time-chat-app-v2/pkg"
)

const address = "0.0.0.0:9090"

func main() {
	chatter := &server.ChatServer{
		RoomUsers: make(map[string][]*server.UserSession),
	}
	mux := http.NewServeMux()
	path, handler := chatv1connect.NewChatServiceHandler(chatter)
	mux.Handle(path, handler)
	err := http.ListenAndServe(address,
		h2c.NewHandler(mux, &http2.Server{}), // Use h2c so we can serve HTTP/2 without TLS.
	)
	if err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
