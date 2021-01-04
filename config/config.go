package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Get() *viper.Viper {
	options := viper.New()

	options.SetDefault("kafka.group.id", "playbook-artifact-validator")
	options.SetDefault("kafka.bootstrap.servers", "kafka:29092")
	options.SetDefault("kafka.auto.offset.reset", "latest")
	options.SetDefault("kafka.auto.commit.interval.ms", 5000)
	options.SetDefault("kafka.request.required.acks", -1) // -1 == "all"
	options.SetDefault("kafka.message.send.max.retries", 15)
	options.SetDefault("kafka.retry.backoff.ms", 100)

	options.SetDefault("topic.request", "platform.upload.playbook")
	options.SetDefault("topic.response", "platform.upload.validation")

	options.SetDefault("metrics.port", 8080)
	options.SetDefault("metrics.path", "/metrics")

	options.SetDefault("openshift.build.commit", "unknown")

	options.SetDefault("log.level", "info")

	options.AutomaticEnv()
	options.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return options
}
