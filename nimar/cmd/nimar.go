package cmd

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
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
	html, err := fs.Sub(htmlFS, "html")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	router.PathPrefix("/nimar/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") //TODO
	})
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") //TODO
		http.FileServer(http.FS(html)).ServeHTTP(w, r)
	})

	//err = http.ListenAndServeTLS(":2222", certFileName, keyFileName, router)
	//if err != nil {
	//panic(err)
	//}
	//_, _ = certFileName, keyFileName
	//_ = wrappedServer
	err = http.ListenAndServe(":2222", router)
	if err != nil {
		panic(err)
	}
}

func Execute() error {
	return cmd.Execute()
}
