package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"playbook-artifact-validator/config"
	"playbook-artifact-validator/ingress"
	probes "playbook-artifact-validator/instrumentation"
	"playbook-artifact-validator/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	cfg   = config.Get()
	topic = cfg.GetString("topic.response")
	log   = utils.GetLoggerOrDie()
)

func OnMessage(msg *kafka.Message) *kafka.Message {
	request := &ingress.Request{}
	err := json.Unmarshal(msg.Value, request)

	if err != nil {
		probes.UnmarshallingError(err)
		return nil
	}

	res, err := http.Get(request.URL)

	if err != nil {
		probes.FetchArchiveError(request, err)
		return nil
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	log.Debugw("Processing request", "account", request.Account, "reqId", request.RequestID)
	response := validateArtifacts(request, data)

	if marshalled, err := json.Marshal(response); err == nil {
		return &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          marshalled,
		}
	} else {
		panic(err) // should never happen
	}
}
