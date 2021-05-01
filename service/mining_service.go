package service

import "luxot.tech/stratum/dto"

type MiningService interface {

	Authorize(request dto.AuthorizeRequest) (bool, error)

	Subscribe() (dto.SubscribeResponse, error)

	Notify() (dto.NotifyResponse, error)

}
