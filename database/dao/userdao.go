package dao

import "go-track/pojo"

type UserDao interface {
	GetByNameAndPasswd(username string, password string) (*pojo.User, error)
}
