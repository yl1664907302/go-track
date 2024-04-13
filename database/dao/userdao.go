package dao

import "kube-auto/pojo"

type UserDao interface {
	GetByNameAndPasswd(username string, password string) (*pojo.User, error)
}
