package dto

type SubscribeResponse struct {
	Extranonce1				string
	Extranonce2_size 		int
	Subscriptions 			[]Subscription
}

type Subscription struct {
	SubscriptionType, Id string
}
