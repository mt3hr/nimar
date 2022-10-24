package cmd

import (
	"embed"
	"encoding/json"
	"io/fs"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
			NimaR()
		},
	}
)

func init() {
	cobra.MousetrapHelpText = ""
	cmd.AddCommand(serverCmd)
}

func NimaR() {
	html, err := fs.Sub(htmlFS, "html")
	if err != nil {
		panic(err)
	}

	server := newServer()
	router := mux.NewRouter()

	updateTablInfos := func() {
		for _, ws := range server.tableListWs {
			b, err := json.Marshal(server.tables)
			if err != nil {
				panic(err)
			}
			ws.Write(b)
		}
	}

	router.PathPrefix("/nimar/ws_list_table").Handler(websocket.Handler(func(ws *websocket.Conn) {
		server.tableListWs = append(server.tableListWs, ws)
		updateTablInfos()
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}))

	router.PathPrefix("/nimar/ws_game_table").Handler(websocket.Handler(func(ws *websocket.Conn) {
		ws.Request().ParseForm()
		tableID := ws.Request().FormValue("tableid")
		playerName := ws.Request().FormValue("playername")
		playerID := ws.Request().FormValue("playerid")

		player := mahjong.NewPlayer(playerName, playerID)
		player.GameTableWs = ws
		if server.tables[tableID].Player1 == nil {
			server.tables[tableID].Player1 = player
		} else if server.tables[tableID].Player2 == nil {
			server.tables[tableID].Player2 = player
		}

		updateTablInfos()
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}))

	router.PathPrefix("/nimar/ws_operators").Handler(websocket.Handler(func(ws *websocket.Conn) {
		ws.Request().ParseForm()
		tableID := ws.Request().FormValue("tableid")
		playerID := ws.Request().FormValue("playerid")

		server.tables[tableID].GetPlayerByID(playerID).OperatorWs = ws
		updateTablInfos()
		wg := sync.WaitGroup{}
		wg.Add(1)
		wg.Wait()
	}))

	router.PathPrefix("/nimar/create_table").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseForm()

		tableName := r.FormValue("tablename")
		id := uuid.New().String()

		server.tables[id] = mahjong.NewTable(id, tableName)
		table := server.tables[id]
		table.GameManager.StartGame()

		b, err := json.Marshal(table)
		if err != nil {
			panic(err)
		}
		w.Write(b)
		updateTablInfos()
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
		b, err := json.Marshal(id)
		if err != nil {
			panic(err)
		}
		w.Write(b)
		updateTablInfos()
	})

	router.PathPrefix("/nimar/execute_operator").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		operator := &mahjong.Operator{}
		err := json.NewDecoder(r.Body).Decode(operator)
		if err != nil {
			panic(err)
		}
		err = server.tables[operator.RoomID].GameManager.ExecuteOperator(operator)
		if err != nil {
			panic(err)
		}
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

func Execute() error {
	return cmd.Execute()
}

type server struct {
	tables      map[string]*mahjong.Table
	tableListWs []*websocket.Conn
	players     map[string]string
}

func newServer() *server {
	return &server{
		tables:  map[string]*mahjong.Table{},
		players: map[string]string{},
	}
}
