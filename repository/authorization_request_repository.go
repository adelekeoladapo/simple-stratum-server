package repository

import "luxot.tech/stratum/model"

type AuthorizationRequestRepository interface {

	SaveAuthorizationRequest(request model.AuthorizationRequest) (model.AuthorizationRequest, error)

	FindAuthorizationRequestById(id int) (model.AuthorizationRequest, error)

	FindAuthorizationRequestByUsername(username string) (model.AuthorizationRequest, error)

}
