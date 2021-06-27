package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func init() {
	// init config (viper)
	initConfig()
	// init logger (logrus)
	initLogger()
}

func main() {
	launchApp()
}

func launchApp() {
	httpAddress := viper.GetString("server.address")

	router := mux.NewRouter()

	register(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         httpAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// serve http
	log.Println("Serve http ON", httpAddress)
	log.Fatal(srv.ListenAndServe())
}
