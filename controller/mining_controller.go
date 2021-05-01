package controller

import "luxot.tech/stratum/dto"

type Mining int

func (mining *Mining) Authorize(args *dto.AuthorizeRequest, result *bool) (e error) {
	*result = true
	return
}

func (mining *Mining) Subscribe() (e error) {
	return
}

func (mining *Mining) Notify() (e error) {
	return
}
