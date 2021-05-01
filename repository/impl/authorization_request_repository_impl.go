package impl

import (
	"luxot.tech/stratum/db"
	"luxot.tech/stratum/model"
)

type AuthorizationRequestRepositoryImpl struct {}

func (impl *AuthorizationRequestRepositoryImpl) SaveAuthorizationRequest(request model.AuthorizationRequest) (res model.AuthorizationRequest, e error) {
	if request.Id > 0 {
		if err := db.Db.Save(&request).Error; err != nil {
			e = err
			return
		}
	} else {
		if err := db.Db.Create(&request).Error; err != nil {
			e = err
			return
		}
	}
	res = request
	return
}

func (impl *AuthorizationRequestRepositoryImpl) FindAuthorizationRequestById(id int) (res model.AuthorizationRequest, e error) {
	if err := db.Db.First(&res, id).Error; err != nil {
		e = err
		return
	}
	return
}

func (impl *AuthorizationRequestRepositoryImpl) FindAuthorizationRequestByUsername(username string) (res model.AuthorizationRequest, e error) {
	if err := db.Db.Where("username = ?", username).First(&res).Error; err != nil {
		e = err
		return
	}
	return
}
