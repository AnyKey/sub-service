package usecase

import (
	"github.com/AnyKey/sub-service/user"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type userUseCase struct {
	GrpcDelivery user.GrpcDelivery
}

func New(ug user.GrpcDelivery) user.Usecase {
	return &userUseCase{
		GrpcDelivery: ug,
	}
}

func (uuc *userUseCase) Token(token string) (*string, error) {

	res, err := uuc.GrpcDelivery.GetToken(token)
	if err != nil {
		log.Errorln("[GetToken] Error:", err)
		return nil, errors.Wrap(err, "[GetToken] Error #1")
	}

	return res, nil
}
