package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"connectrpc.com/connect"
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
	mux.HandleFunc("/", StaticHandler)

	interceptors := connect.WithInterceptors(server.NewAuthInterceptor())
	path, handler := chatv1connect.NewChatServiceHandler(chatter, interceptors)
	mux.Handle(path, handler)
	err := http.ListenAndServe(address,
		h2c.NewHandler(mux, &http2.Server{}), // Use h2c so we can serve HTTP/2 without TLS.
	)
	if err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

//go:embed ui/build
var files embed.FS

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	subpath, _ := fs.Sub(files, "ui/build")
	fs := http.FileServer(http.FS(subpath))
	if r.URL.Path == "/" {
		fs.ServeHTTP(w, r)
		return
	}
	f, err := subpath.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
	if err == nil {
		defer f.Close()
	}
	if os.IsNotExist(err) {
		r.URL.Path = "/"
	}
	fs.ServeHTTP(w, r)
}
