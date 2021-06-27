package main

import (
	"github.com/AnyKey/sub-service/client"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	userGrpcDelivery "github.com/AnyKey/sub-service/user/delivery/grpc"
	userHttpDelivery "github.com/AnyKey/sub-service/user/delivery/http"
	userUseCase "github.com/AnyKey/sub-service/user/usecase"
)

// register usecase, delivery, repository for each entity
func register(router *mux.Router, s *grpc.Server) {
	emailRegister(router, s)
	client.Template(router)
}

// register user entity
func emailRegister(router *mux.Router, s *grpc.Server) {
	HttpDelivery := userHttpDelivery.New(router)

	userUCase := userUseCase.New(HttpDelivery)

	userGrpcDelivery.Launch(s, userUCase)

}
