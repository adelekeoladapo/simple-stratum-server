package impl

import (
	"luxot.tech/stratum/db"
	"luxot.tech/stratum/model"
)

type SubscriptionRepositoryImpl struct{}

func (impl *SubscriptionRepositoryImpl) SaveSubscription(subscription model.Subscription) (res model.Subscription, e error) {
	if subscription.Id > 0 {
		if err := db.Db.Save(&subscription).Error; err != nil {
			e = err
			return
		}
	} else {
		if err := db.Db.Create(&subscription).Error; err != nil {
			e = err
			return
		}
	}
	res = subscription
	return
}
