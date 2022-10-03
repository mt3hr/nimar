package main

import (
	"net/http"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/mt3hr/nimar"
	"google.golang.org/grpc"
)

func main() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	certFileName = filepath.Join(homeDir, "cert.pem")
	keyFileName = filepath.Join(homeDir, "key.pem")

	nimarServer, err := nimar.NewNimeRServer()
	server := grpc.NewServer()
	nimar.RegisterNimaRServer(server, nimarServer)

	http.Handle("/nimar/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	}))

	err = http.ListenAndServeTLS(":9999", certFileName, keyFileName, nil)
	if err != nil {
		panic(err)
	}
}
