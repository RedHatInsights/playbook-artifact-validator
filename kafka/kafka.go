package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md

func NewProducer(config *viper.Viper) (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        config.GetString("kafka.bootstrap.servers"),
		"request.required.acks":    config.GetInt("kafka.request.required.acks"),
		"message.send.max.retries": config.GetInt("kafka.message.send.max.retries"),
		"retry.backoff.ms":         config.GetInt("kafka.retry.backoff.ms"),
	})
}

func NewConsumer(config *viper.Viper) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       config.GetString("kafka.bootstrap.servers"),
		"group.id":                config.GetString("kafka.group.id"),
		"auto.offset.reset":       config.GetString("kafka.auto.offset.reset"),
		"auto.commit.interval.ms": config.GetInt("kafka.auto.commit.interval.ms"),
	})

	if err != nil {
		return nil, err
	}

	consumer.SubscribeTopics([]string{config.GetString("topic.request")}, nil)
	return consumer, nil
}
