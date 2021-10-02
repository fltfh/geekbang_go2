package service

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"template/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	logger *logrus.Logger
	user   *biz.UserService
}
