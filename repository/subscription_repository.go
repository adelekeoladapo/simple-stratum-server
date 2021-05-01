package repository

import "luxot.tech/stratum/model"

type SubscriptionRepository interface {

	SaveSubscription(subscription model.Subscription) (model.Subscription, error)

}
