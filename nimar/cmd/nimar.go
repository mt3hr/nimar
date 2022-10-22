package cmd

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
	"github.com/mt3hr/nimar/mahjong"
	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
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

	server := newServer()
	router := mux.NewRouter()

	a := func() {
		tables := []*TableInfo{}
		for _, table := range server.tables {
			playerNames := []string{}
			if table.GetPlayer1() != nil {
				playerNames = append(playerNames, table.GetPlayer1().GetName())
			}
			if table.GetPlayer2() != nil {
				playerNames = append(playerNames, table.GetPlayer2().GetName())
			}
			tables = append(tables, &TableInfo{
				TableID:     table.GetID(),
				TableName:   table.GetName(),
				PlayerNames: playerNames,
			})
		}

		for _, ws := range server.tableListWs {
			b, err := json.Marshal(tables)
			if err != nil {
				panic(err)
			}
			ws.Write(b)
		}
	}

	router.PathPrefix("/nimar/ws_list_table").Handler(websocket.Handler(func(ws *websocket.Conn) {
		server.tableListWs = append(server.tableListWs, ws)
		wg := sync.WaitGroup{}
		wg.Add(1)
		a()
		wg.Wait()
	}))

	router.PathPrefix("/nimar/ws_game_table").Handler(websocket.Handler(func(ws *websocket.Conn) {
		ws.Request().ParseForm()
		roomID := ws.Request().FormValue("roomid")
		server.gameTableWs[roomID] = append(server.gameTableWs[roomID], ws)
		wg := sync.WaitGroup{}
		wg.Add(1)
		a()
		wg.Wait()
	}))

	router.PathPrefix("/nimar/create_table").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseForm()
		tableName := r.FormValue("table_name")
		table := &TableInfo{
			TableName: tableName,
			TableID:   uuid.New().String(),
		}
		server.tables[table.TableID] = mahjong.NewTable(table.TableID, table.TableName)

		b, err := json.Marshal(table)
		if err != nil {
			panic(err)
		}
		w.Write(b)
		a()
	})

	router.PathPrefix("/nimar/get_player_id").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ipAddress := r.Header.Get("X-Forwarded-For")
		id := ""
		ok := false
		if id, ok = server.players[ipAddress]; !ok {
			id = uuid.New().String()
			server.players[ipAddress] = id
		}
		playerID := &PlayerIDInfo{
			PlayerID: id,
		}
		b, err := json.Marshal(playerID)
		if err != nil {
			panic(err)
		}
		w.Write(b)
		a()
	})

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.FileServer(http.FS(html)).ServeHTTP(w, r)
	})

	err = http.ListenAndServe(":2222", router)
	if err != nil {
		panic(err)
	}
}

type PlayerIDInfo struct {
	PlayerID string `json:"player_id"`
}

type TableInfo struct {
	TableID     string   `json:"table_id"`
	TableName   string   `json:"table_name"`
	PlayerNames []string `json:"player_names"`
}

func Execute() error {
	return cmd.Execute()
}

type server struct {
	tables      map[string]*mahjong.Table
	tableListWs []*websocket.Conn
	players     map[string]string
	gameTableWs map[string][]*websocket.Conn
}

func newServer() *server {
	return &server{
		tables:      map[string]*mahjong.Table{},
		players:     map[string]string{},
		gameTableWs: map[string][]*websocket.Conn{},
	}
}
