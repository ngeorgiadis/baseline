package main

import (
	"fmt"
	"mime"
	"net/http"

	"os"
	"os/signal"
	"syscall"

	"github.com/ngeorgiadis/baseline/cmd/server/config"
	"github.com/ngeorgiadis/baseline/cmd/server/handlers"
	log "github.com/sirupsen/logrus"
)

var appConfig *config.Config

func init() {

	var err error
	appConfig, err = config.New("settings.cfg")
	if err != nil {
		panic(err.Error())
	}

	mime.AddExtensionType(".js", "text/javascript; charset=utf-8")
}

func main() {

	r, err := handlers.GetRouter(appConfig)
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", appConfig.App.Port),
		Handler: r,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT)

	errorChan := make(chan error, 1)
	go func() {
		log.Infof("starting service at address: [ %v ]", server.Addr)
		errorChan <- server.ListenAndServe()
	}()

	select {
	case err := <-errorChan:
		log.Fatal(err.Error())
	case sig := <-shutdown:
		log.Infof("interupted by user: [ %v ] ", sig.String())
		server.Close()
	}
}
