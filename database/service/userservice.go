package service

import "kube-auto/pojo"

type UserService interface {
	GetUserByNameAndPasswd(username string, password string) (*pojo.User, error)
}
