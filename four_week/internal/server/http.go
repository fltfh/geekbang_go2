package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"template/internal/biz"
	"template/internal/config"
	"template/pkg/manager"
)

type server struct {
	g      *gin.Engine
	srv    *http.Server
	logger *logrus.Logger
}

func (s *server) Start(ctx context.Context) error {
	err := s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func NewHTTPServer(logger *logrus.Logger, info config.AppInfo, userSrv *biz.UserService) manager.Server {

	g := gin.Default()
	g.Use(gin.LoggerWithWriter(logger.Out))

	g.GET("/ping", func(c *gin.Context) {

		data, _ := userSrv.GetById(context.Background(), 10)
		data1, _ := json.Marshal(data)
		c.JSONP(http.StatusOK, gin.H{
			"code": 0,
			"message": "调用成功",
			"data": string(data1),
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", info.Host, info.Port),
		Handler: g,
	}
	logger.Printf("开始启动http服务：%s:%s\n", info.Host, info.Port)
	return &server{
		g:      gin.Default(),
		srv:    srv,
		logger: logger,
	}
}
