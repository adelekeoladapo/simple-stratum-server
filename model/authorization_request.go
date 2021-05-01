package model

import "time"

type AuthorizationRequest struct {
	Id 						int
	Username 				string
	LastLoggedInTime 		time.Time
}
