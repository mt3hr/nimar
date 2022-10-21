package cmd

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
	"github.com/mt3hr/nimar"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (

	//go:embed html
	htmlFS embed.FS // htmlファイル郡

	cmd = &cobra.Command{
		Run: func(_ *cobra.Command, _ []string) {
		},
	}

	serverCmd = &cobra.Command{
		Use: "server",
		Run: func(_ *cobra.Command, _ []string) {
			NimaR(certFileName, keyFileName)
		},
	}
)

var (
	certFileName string
	keyFileName  string
)

func init() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	certFileName = filepath.Join(homeDir, "cert.pem")
	keyFileName = filepath.Join(homeDir, "key.pem")

	cobra.MousetrapHelpText = ""
	cmd.AddCommand(serverCmd)
	cmd.Flags().StringVarP(&certFileName, "cert", "c", certFileName, "")
	cmd.Flags().StringVarP(&keyFileName, "key", "k", keyFileName, "")
}

func NimaR(certFileName, keyFileName string) {
	nimarServer, err := nimar.NewNimaRServer()
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	nimar.RegisterNimaRServer(server, nimarServer)
	reflection.Register(server)

	html, err := fs.Sub(htmlFS, "html")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.PathPrefix("/NimaR/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			server.ServeHTTP(w, r)
		})
	router.PathPrefix("/").Handler(http.FileServer(http.FS(html)))

	err = http.ListenAndServeTLS(":2222", certFileName, keyFileName, router)
	// _, _ = certFileName, keyFileName
	///err = http.ListenAndServe(":2222", router)
	if err != nil {
		panic(err)
	}
}

func Execute() error {
	return cmd.Execute()
}
