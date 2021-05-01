package impl

import (
	"luxot.tech/stratum/dto"
	"luxot.tech/stratum/model"
	"luxot.tech/stratum/repository/impl"
	"luxot.tech/stratum/utils"
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
	var subscription model.Subscription = model.Subscription{
		SubscriptionId1: utils.GenerateRandomString(20),
		SubscriptionId2: utils.GenerateRandomString(20),
		Extranonce1: utils.GenerateRandomString(10),
		DateCreated: time.Now(),
	}
	_, err := impl.subscriptionRepository.SaveSubscription(subscription)
	if err != nil {
		e = err
		return
	}
	response = dto.SubscribeResponse{
		Extranonce1: subscription.Extranonce1,
		Extranonce2_size: 4,
		Subscriptions: []dto.Subscription{
			{
				"mining.set_difficulty",
				subscription.SubscriptionId1,
			},
			{
				"mining.notify",
				subscription.SubscriptionId2,
			},
		},
	}
	return
}

func (impl MiningServiceImpl) Notify() (response dto.NotifyResponse, e error) {
	return
}
