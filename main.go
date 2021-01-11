package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"playbook-artifact-validator/config"
	"playbook-artifact-validator/handler"
	instrumentation "playbook-artifact-validator/instrumentation"
	"playbook-artifact-validator/utils"
	"syscall"
	"time"

	"playbook-artifact-validator/kafka"

	k "github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	kafkaTimeout = 5000
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	log := utils.GetLoggerOrDie()

	config := config.Get()

	consumer, err := kafka.NewConsumer(config)
	utils.DieOnError(err)

	producer, err := kafka.NewProducer(config)
	utils.DieOnError(err)

	topic := config.GetString("topic.response")

	mux := http.NewServeMux()
	instrumentation.StartMetrics(config, mux)

	mux.HandleFunc("/ready", func(resp http.ResponseWriter, _ *http.Request) {
		if _, err := producer.GetMetadata(&topic, false, kafkaTimeout); err != nil {
			resp.WriteHeader(http.StatusServiceUnavailable)
		} else {
			resp.WriteHeader(http.StatusOK)
		}
	})

	mux.HandleFunc("/live", func(resp http.ResponseWriter, _ *http.Request) {
		resp.WriteHeader(http.StatusOK)
	})

	log.Infof("Listening on port %d", config.GetInt("metrics.port"))
	go http.ListenAndServe(fmt.Sprintf(":%d", config.GetInt("metrics.port")), mux)

	defer func() {
		log.Info("Shutting down")
		producer.Flush(kafkaTimeout)
		producer.Close()
		consumer.Close()

		log.Info("Shutdown complete")
		log.Sync()
	}()

	log.Infow("Validator started", "version", config.GetString("openshift.build.commit"))

	for {
		if signal := utils.PeekSignalChannel(signals); signal != nil {
			log.Infow("Received signal", "signal", signal)
			break
		}

		msg, err := consumer.ReadMessage(1 * time.Second)

		if err != nil {
			if err.(k.Error).Code() == k.ErrTimedOut {
				continue
			}

			log.Errorw("Consumer error", "error", err)
			break
		}

		if response := handler.OnMessage(msg); response != nil {
			if err = producer.Produce(response, nil); err != nil {
				log.Errorw("Producer error", "error", err)
				break
			}
		}
	}
}
