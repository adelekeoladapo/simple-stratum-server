package impl

import (
	"luxot.tech/stratum/dto"
	"luxot.tech/stratum/repository/impl"
)

type MiningServiceImpl struct {
	authorizationRequestRepository impl.AuthorizationRequestRepositoryImpl
	subscriptionRepository impl.SubscriptionRepositoryImpl
}

func (impl MiningServiceImpl) Authorize(request dto.AuthorizeRequest) (response bool, e error) {
	return
}

func (impl MiningServiceImpl) Subscribe() (response dto.SubscribeResponse, e error) {
	return
}

func (impl MiningServiceImpl) Notify() (response dto.NotifyResponse, e error) {
	return
}
