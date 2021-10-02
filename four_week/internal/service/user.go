package service

import (
	"github.com/sirupsen/logrus"
	"template/internal/biz"
)

func NewUserService(user *biz.UserService, logger *logrus.Logger) *UserService {
	return &UserService{
		user: user,
		logger:  logger,
	}
}