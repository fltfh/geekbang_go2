package manager

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type App struct {
	opts options
	ctx context.Context
	cancel context.CancelFunc
}


func New(logger *logrus.Logger, opts ...Option) *App {
	o := options{
		ctx:              context.Background(),
		logger:           logger,
		sigs:             []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}



func (a * App) Run() error{
	// 初始化errgroup
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	// 启动option中的server
	for _, srv := range a.opts.servers {
		eg.Go(func() error {
			<- ctx.Done()
			return srv.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}
	wg.Wait()
	// 等待所有server启动并开启信号检测
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				err := a.Stop()
				if err != nil {
					a.opts.logger.Errorf("failed to stop app: %v", err)
				}
			}
		}
	})

	if err := eg.Wait(); err != nil && ! errors.Is(err, context.Canceled){
		return err
	}
	return nil
}

func (a* App) Stop() error {
	a.opts.logger.Printf("服务开始停止...")
	a.cancel()
	return nil
}
