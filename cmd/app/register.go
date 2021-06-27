package main

import (
	"github.com/AnyKey/sub-service/client"
	userGrpcDelivery "github.com/AnyKey/sub-service/user/delivery/grpc"
	userHttpDelivery "github.com/AnyKey/sub-service/user/delivery/http"
	userUseCase "github.com/AnyKey/sub-service/user/usecase"
	"github.com/gorilla/mux"
)

// register usecase, delivery, repository for each entity
func register(router *mux.Router) {
	emailRegister(router)
	client.Template(router)
}

// register user entity
func emailRegister(router *mux.Router) {
	grpcDelivery := userGrpcDelivery.NewClient()

	userUCase := userUseCase.New(grpcDelivery)

	go userHttpDelivery.Launch(router, userUCase)

}
