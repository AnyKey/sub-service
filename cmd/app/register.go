package main

import (
	"database/sql"
	"github.com/AnyKey/sub-service/client"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	userGrpcDelivery "github.com/AnyKey/sub-service/user/delivery/grpc"
	userHttpDelivery "github.com/AnyKey/sub-service/user/delivery/http"
	userDBRepository "github.com/AnyKey/sub-service/user/repository/postgresql"
	userUseCase "github.com/AnyKey/sub-service/user/usecase"
)

// register usecase, delivery, repository for each entity
func register(router *mux.Router, s *grpc.Server, db *sql.DB) {
	emailRegister(router, s, db)
	client.Template(router)
}

// register user entity
func emailRegister(router *mux.Router, s *grpc.Server, db *sql.DB) {
	DBRepo := userDBRepository.New(db)
	HttpDelivery := userHttpDelivery.New(router)

	userUCase := userUseCase.New(DBRepo, HttpDelivery)

	userGrpcDelivery.Launch(s, userUCase)

}
