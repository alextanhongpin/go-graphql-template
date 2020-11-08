package app

import (
	"context"
	stdlog "log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/alextanhongpin/go-graphql-template/pkg/logger"
	"github.com/alextanhongpin/pkg/grace"
)

type LoggerConfig struct {
	Name string `envconfig:"NAME" default:"go-backend"`
	Env  string `envconfig:"ENV" default:"development"`
}

func NewLogger(sg *grace.ShutdownGroup) *zap.Logger {
	var cfg LoggerConfig
	if err := envconfig.Process("", &cfg); err != nil {
		stdlog.Fatal(err)
	}

	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	log := logger.New(cfg.Env,
		zap.String("app", cfg.Name),
		zap.String("host", host))

	zap.ReplaceGlobals(log)
	sg.Add(func(_ context.Context) {
		if err := log.Sync(); err != nil {
			stdlog.Println(err)
		}
	})

	return log
}
