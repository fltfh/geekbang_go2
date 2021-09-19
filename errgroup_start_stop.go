package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"html"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	handler := &Handler{}
	return &Server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	fmt.Println("service start stop ..aaaaaaaaaaaa")
	return s.srv.Shutdown(context.Background())
}

type App struct {
	servers []Server
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewApp(servers []Server) *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		servers: servers,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (a *App) Start() error {
	eg, ctx := errgroup.WithContext(a.ctx)

	wg := sync.WaitGroup{}
	for _, s := range a.servers {
		// 等待终端信号
		eg.Go(func() error {
			<-ctx.Done()
			return s.Stop()
		})
		wg.Add(1)
		// 服务启动
		eg.Go(func() error {
			time.Sleep(5 * time.Second)
			wg.Done()

			return s.Start()
		})
	}

	// 等待所有服务启动后再开始注册信号
	fmt.Println("all server is started ...................")
	wg.Wait()

	// 注册信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	fmt.Println("register signal success ...................")

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				fmt.Println("stop app ...................")
				return a.Stop()
			}
		}
	})

	if err := eg.Wait(); err != nil || !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	a.cancel()
	return nil
}

func main() {
	srv := NewServer()
	app := NewApp([]Server{*srv})
	app.Start()
}
