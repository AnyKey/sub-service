package main

import (
	"database/sql"
	"google.golang.org/grpc"

	//emailGrpcDelivery "git.zuzex.com/bots/back/email.git/email/delivery/grpc"
	//emailSmtpDelivery "git.zuzex.com/bots/back/email.git/email/delivery/smtp"
	//emailDBRepository "git.zuzex.com/bots/back/email.git/email/repository/postgresql"
	//emailUsecase "git.zuzex.com/bots/back/email.git/email/usecase"
)

// register usecase, delivery, repository for each entity
func register(s *grpc.Server, db *sql.DB) {
	emailRegister(s, db)
}

// register email entity
func emailRegister(s *grpc.Server, db *sql.DB) {
	emailDBRepo := emailDBRepository.New(db)
	emailSMTPDelivery := emailSmtpDelivery.Delivery{}

	emailUcase := emailUsecase.New(emailDBRepo, emailSMTPDelivery)

	go emailNatsDelivery.NewConsumer(emailUcase).StartWatch(jsc)

	emailGrpcDelivery.Launch(s, emailUcase)

}
