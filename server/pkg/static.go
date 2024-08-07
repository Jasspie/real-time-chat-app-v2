package server

import (
	"net/http"
	"os"
	"path"
	"strings"
)

const FSPATH = "ui/build"

func HandleStatic(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/" {
		fullPath := FSPATH + strings.TrimPrefix(path.Clean(req.URL.Path), "/")
		_, err := os.Stat(fullPath)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
			// Requested file does not exist so we return the default (resolves to index.html)
			req.URL.Path = "/"
		}
	}
	fs := http.FileServer(http.Dir(FSPATH))
	fs.ServeHTTP(w, req)
}
