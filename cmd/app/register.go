package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"sub-github.com/AnyKey/sub-service.git/client"

	userGrpcDelivery "sub-github.com/AnyKey/sub-service.git/user/delivery/grpc"
	userHttpDelivery "sub-github.com/AnyKey/sub-service.git/user/delivery/http"
	userDBRepository "sub-github.com/AnyKey/sub-service.git/user/repository/postgresql"
	userUseCase "sub-github.com/AnyKey/sub-service.git/user/usecase"
)

// register usecase, delivery, repository for each entity
func register(router *mux.Router,s *grpc.Server, db *sql.DB) {
	emailRegister(router, s, db)
	client.Template(router)
}

// register user entity
func emailRegister(router *mux.Router,s *grpc.Server, db *sql.DB) {
	DBRepo := userDBRepository.New(db)
	HttpDelivery := userHttpDelivery.New(router)

	userUCase := userUseCase.New(DBRepo, HttpDelivery)

	userGrpcDelivery.Launch(s, userUCase)

}
