package service

import "go-track/pojo"

type UserService interface {
	GetUserByNameAndPasswd(username string, password string) (*pojo.User, error)
}
