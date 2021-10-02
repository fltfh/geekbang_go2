// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"template/internal/biz"
	"template/internal/config"
	"template/internal/repo"
	"template/internal/server"
	"template/internal/service"
	"template/pkg/manager"
)

func initApp(logger *logrus.Logger, db *gorm.DB, info config.AppInfo) (*manager.App, error) {
	panic(wire.Build(repo.ProviderSet, biz.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
