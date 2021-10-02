package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/mysql"
	"template/internal/config"
	"template/pkg/manager"
)

var (

	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	// 初始化参数
	flag.Parse()
	// 初始化配置文件
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	// 初始化日志
	logger := logrus.New()

	// 初始化数据库
	db, err := gorm.Open("mysql", config.Config.DbDsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app, err := initApp(logger, db, config.Config.App)
	if err != nil {
		panic(err)
	}
	//// 启动app
	if err := app.Run(); err != nil {
		panic(err)
	}
}


func newApp(logger *logrus.Logger, hs manager.Server, info config.AppInfo) *manager.App {
	return manager.New(
		logger,
		manager.Name(info.Name),
		manager.Version(info.Version),
		manager.Servers(
			hs,
		),
	)
}
