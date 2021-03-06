package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"playbook-artifact-validator/config"
	"playbook-artifact-validator/ingress"
	probes "playbook-artifact-validator/instrumentation"
	"playbook-artifact-validator/utils"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	cfg    = config.Get()
	topic  = cfg.GetString("topic.response")
	log    = utils.GetLoggerOrDie()
	client = &http.Client{
		Timeout: time.Duration(cfg.GetInt64("storage.timeout") * int64(time.Second)),
	}
)

func OnMessage(msg *kafka.Message) *kafka.Message {
	request := &ingress.Request{}
	err := json.Unmarshal(msg.Value, request)

	if err != nil {
		probes.UnmarshallingError(err)
		return nil
	}

	response := onRequest(request)

	if marshalled, err := json.Marshal(response); err == nil {
		return &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          marshalled,
		}
	} else {
		panic(err) // should never happen
	}
}

func onRequest(request *ingress.Request) *ingress.Response {
	if request.Size > cfg.GetInt64("max.size") {
		probes.FileTooLarge(request)

		return ingress.NewResponse(request, "failure")
	}

	res, err := utils.DoGetWithRetry(client, request.URL, cfg.GetInt("storage.retries"))

	if err != nil {
		probes.FetchArchiveError(request, err)
		return nil
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	log.Debugw("Processing request", "account", request.Account, "reqId", request.RequestID)
	return validateArtifacts(request, data)
}
