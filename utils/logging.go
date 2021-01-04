package utils

import (
	"playbook-artifact-validator/config"

	"go.uber.org/zap"
)

func GetLoggerOrDie() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.Level.UnmarshalText([]byte(config.Get().GetString("log.level")))
	log, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	return log.Sugar()
}
