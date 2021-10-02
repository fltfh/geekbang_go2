package manager

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type Option func(o *options)

type options struct {
	id      string
	name    string
	version string
	ctx     context.Context
	sigs    []os.Signal
	logger  *logrus.Logger
	servers []Server
}

// ID with biz id.
func ID(id string) Option {
	return func(o *options) { o.id = id }
}

// Name with biz name.
func Name(name string) Option {
	return func(o *options) { o.name = name }
}

// Version with biz version.
func Version(version string) Option {
	return func(o *options) { o.version = version }
}

// Logger with biz logger.
func Logger(logger *logrus.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Servers with transport servers.
func Servers(srv ...Server) Option {
	return func(o *options) { o.servers = srv }
}
