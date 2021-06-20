package main

import (
	"net"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	var (
		grpcP       = viper.GetString("grpc.port")
		httpAddress = viper.GetString("server.address")
	)

	// creating singleton grpc server
	s := grpc.NewServer()

	// create database instance
	dbConn, err := initPostgresInstance()
	if err != nil {
		log.Fatalf("connect database error %s", err)
	}
	defer dbConn.Close()

	// echo http server
	e := echo.New()


	// register entities
	register(s, dbConn)

	// creating listener on port for http and grpc
	listener, err := net.Listen("tcp", grpcP)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// serve grpc
	log.Infof("Start serve grpc on %s port", grpcP)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// serve http
	log.Infof("Start serve http on %s port", httpAddress)
	log.Fatalln("ECHO: ", e.Start(httpAddress))
}
