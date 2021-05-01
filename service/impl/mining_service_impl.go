package impl

import (
	"luxot.tech/stratum/dto"
	"luxot.tech/stratum/repository/impl"
	"time"
)

type MiningServiceImpl struct {
	authorizationRequestRepository impl.AuthorizationRequestRepositoryImpl
	subscriptionRepository impl.SubscriptionRepositoryImpl
}

func (impl MiningServiceImpl) Authorize(request dto.AuthorizeRequest) (response bool, e error) {
	authReq, _ := impl.authorizationRequestRepository.FindAuthorizationRequestByUsername(request.Username)
	authReq.Username = request.Username
	authReq.LastLoggedInTime = time.Now()
	_, err := impl.authorizationRequestRepository.SaveAuthorizationRequest(authReq)
	if err != nil {
		e = err
		return
	}
	response = true
	return
}

func (impl MiningServiceImpl) Subscribe() (response dto.SubscribeResponse, e error) {
	return
}

func (impl MiningServiceImpl) Notify() (response dto.NotifyResponse, e error) {
	return
}
