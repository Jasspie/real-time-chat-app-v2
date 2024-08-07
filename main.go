package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	chatv1connect "github.com/Jasspie/real-time-chat-app-v2/server/chat/v1/v1connect"
	server "github.com/Jasspie/real-time-chat-app-v2/server/pkg"
)

const address = "localhost:8030"

func main() {
	chatter := &server.ChatServer{
		RoomUsers: make(map[string][]*server.UserSession),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/login", server.RootHandler)
	mux.HandleFunc("/google_callback", server.CallbackHandler)
	mux.HandleFunc("/", server.HandleStatic)

	// interceptors := connect.WithInterceptors(server.NewAuthInterceptor())
	path, handler := chatv1connect.NewChatServiceHandler(chatter)
	mux.Handle(path, handler)
	err := http.ListenAndServe(address,
		h2c.NewHandler(mux, &http2.Server{}), // Use h2c so we can serve HTTP/2 without TLS.
	)
	if err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
