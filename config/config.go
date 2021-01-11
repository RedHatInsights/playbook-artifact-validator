package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
)

func Get() *viper.Viper {
	options := viper.New()

	options.SetDefault("kafka.group.id", "playbook-artifact-validator")
	options.SetDefault("kafka.auto.offset.reset", "latest")
	options.SetDefault("kafka.auto.commit.interval.ms", 5000)
	options.SetDefault("kafka.request.required.acks", -1) // -1 == "all"
	options.SetDefault("kafka.message.send.max.retries", 15)
	options.SetDefault("kafka.retry.backoff.ms", 100)

	options.SetDefault("openshift.build.commit", "unknown")
	options.SetDefault("runner.schema", "./schemas/runner.yaml")

	options.SetDefault("storage.timeout", 10)
	options.SetDefault("storage.retries", 3)

	options.SetDefault("log.level", "debug")

	if os.Getenv("CLOWDER_ENABLED") != "false" {
		options.SetDefault("kafka.bootstrap.servers", strings.Join(clowder.KafkaServers, ","))
		options.SetDefault("topic.request", clowder.KafkaTopics["platform.upload.playbook"].Name)
		options.SetDefault("topic.response", clowder.KafkaTopics["platform.upload.validation"].Name)

		options.SetDefault("metrics.port", clowder.LoadedConfig.MetricsPort)
		options.SetDefault("metrics.path", clowder.LoadedConfig.MetricsPath)

		options.SetDefault("log.cw.accessKeyId", clowder.LoadedConfig.Logging.Cloudwatch.AccessKeyId)
		options.SetDefault("log.cw.secretAccessKey", clowder.LoadedConfig.Logging.Cloudwatch.SecretAccessKey)
		options.SetDefault("log.cw.region", clowder.LoadedConfig.Logging.Cloudwatch.Region)
		options.SetDefault("log.cw.group", clowder.LoadedConfig.Logging.Cloudwatch.LogGroup)
	} else {
		options.SetDefault("kafka.bootstrap.servers", "kafka:29092")
		options.SetDefault("topic.request", "platform.upload.playbook")
		options.SetDefault("topic.response", "platform.upload.validation")

		options.SetDefault("metrics.port", 9001)
		options.SetDefault("metrics.path", "/metrics")
	}

	options.AutomaticEnv()
	options.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return options
}
