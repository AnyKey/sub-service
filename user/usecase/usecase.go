package usecase

import "github.com/AnyKey/sub-service/user"

type userUseCase struct {
	userHttpDelivery user.HttpDelivery
}

func New(uh user.HttpDelivery) user.Usecase {
	return &userUseCase{
		userHttpDelivery: uh,
	}
}

func (uuc *userUseCase) Token() error {
	return nil
}
