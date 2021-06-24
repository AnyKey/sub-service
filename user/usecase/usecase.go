package usecase

import "github.com/AnyKey/sub-service/user"

type userUseCase struct {
	userRepo         user.Repository
	userHttpDelivery user.HttpDelivery
}

func New(ur user.Repository, uh user.HttpDelivery) user.Usecase {
	return &userUseCase{
		userRepo:         ur,
		userHttpDelivery: uh,
	}
}

func (uuc *userUseCase) Token() error {
	return nil
}
