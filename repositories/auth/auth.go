package auth

import "github.com/youtrolledhahaha/youtrolledhahahaXxxentities"

type Repository interface {
	Insert(auth entities.Auth) error
	Update(auth *entities.Auth) error
	GetFirst() (*entities.Auth, error)
}
