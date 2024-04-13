package servicelmpl

import (
	"kube-auto/database/dao"
	"kube-auto/pojo"
)

type UserServiceImpl struct {
	userdao dao.UserDao
}

func NewUserServiceImpl(userdao dao.UserDao) *UserServiceImpl {
	return &UserServiceImpl{userdao: userdao}
}

func (s *UserServiceImpl) GetUserByNameAndPasswd(username string, password string) (*pojo.User, error) {
	return s.userdao.GetByNameAndPasswd(username, password)
}
