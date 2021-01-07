package handler

import (
	"context"
	"io/ioutil"
	"playbook-artifact-validator/config"
	"playbook-artifact-validator/ingress"
	probes "playbook-artifact-validator/instrumentation"
	"playbook-artifact-validator/utils"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/qri-io/jsonschema"
)

var schema *jsonschema.Schema

func init() {
	file, err := ioutil.ReadFile(config.Get().GetString("runner.schema"))
	utils.DieOnError(err)

	schema = &jsonschema.Schema{}

	err = yaml.Unmarshal(file, schema)
	utils.DieOnError(err)
}

func validateArtifacts(request *ingress.Request, data []byte) *ingress.Response {
	var (
		parserError      error
		validationErrors = []jsonschema.KeyError{}
	)

	response := &ingress.Response{
		Request: *request,
	}

	events := strings.Split(string(data), "\n")

	for _, event := range events {
		if len(strings.TrimSpace(event)) == 0 {
			continue
		}

		var errors []jsonschema.KeyError
		errors, parserError = schema.ValidateBytes(context.TODO(), []byte(event))
		if parserError != nil {
			break
		}

		validationErrors = append(validationErrors, errors...)
	}

	switch {
	case parserError != nil:
		response.Validation = "failure"
		probes.ValidationFailed(request, parserError)
	case len(validationErrors) > 0:
		response.Validation = "failure"
		probes.ValidationFailed(request, validationErrors[0])
	default:
		response.Validation = "success"
		probes.ValidationSuccess(request)
	}

	return response
}
