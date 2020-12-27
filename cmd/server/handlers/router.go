package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ngeorgiadis/baseline/cmd/server/config"
	"github.com/ngeorgiadis/baseline/cmd/server/wshub"
)

// var appConfig *config.Config

// var conn *sqlserver.Connection

// GetRouter ...
func GetRouter(appConfig *config.Config) (*mux.Router, error) {

	r := mux.NewRouter()

	hub := wshub.New()
	go hub.Run()

	_, err := config.New("settings.cfg")
	if err != nil {
		return nil, err
	}

	//
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/ui/", 302)
	})

	//ui
	ui := r.PathPrefix("/ui").Subrouter()
	ui.HandleFunc("/", homeHandler)

	// ui.Use(isAuthenticated(appConfig, conn))

	// application
	app := r.PathPrefix("/app").Subrouter()
	app.HandleFunc("/home", homeHandler)
	app.HandleFunc("/list", listHandler(appConfig))
	app.HandleFunc("/file/{filename}", fileHandler(appConfig))
	app.HandleFunc("/graph/{id}", getGraphData2(appConfig, false))
	app.HandleFunc("/graph/{id}/maxcore", getGraphData2(appConfig, true))

	// websocket handler
	ws := r.PathPrefix("/ws").Subrouter()
	ws.HandleFunc("/", serveWs(hub))
	// ws.Use(isAuthenticated(appConfig, conn))

	// static
	r.PathPrefix("/ui/").Handler(http.StripPrefix("/ui/", http.FileServer(http.Dir("../../ui/dist"))))

	// middleware
	// app.Use(isAuthenticated(appConfig, conn))

	return r, nil
}
