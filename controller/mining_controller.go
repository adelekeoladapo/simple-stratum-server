package controller

import (
	"luxot.tech/stratum/dto"
	"luxot.tech/stratum/service"
	"luxot.tech/stratum/service/impl"
)

var miningService service.MiningService = impl.MiningServiceImpl{}

type Mining int

func (mining *Mining) Authorize(args *dto.AuthorizeRequest, result *bool) (e error) {
	*result, e = miningService.Authorize(*args)
	return
}

func (mining *Mining) Subscribe(args *string, result *dto.SubscribeResponse) (e error) {
	*result, e = miningService.Subscribe()
	return
}

func (mining *Mining) Notify(args string, result *dto.NotifyResponse) (e error) {
	*result, e = miningService.Notify()
	return
}
