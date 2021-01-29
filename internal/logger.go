package internal

import (
	stdlog "log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/alextanhongpin/go-graphql-template/pkg/logger"
)

type LoggerConfig struct {
	Name string `envconfig:"NAME" default:"go-backend"`
	Env  string `envconfig:"ENV" default:"development"`
}

func NewLoggerConfig() LoggerConfig {
	var cfg LoggerConfig
	if err := envconfig.Process("", &cfg); err != nil {
		stdlog.Fatal(err)
	}
	return cfg
}

func NewLogger(cfg LoggerConfig) (*zap.Logger, func()) {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	log := logger.New(cfg.Env,
		zap.String("app", cfg.Name),
		zap.String("host", host))

	zap.ReplaceGlobals(log)

	return log, func() {
		_ = log.Sync()
	}
}
