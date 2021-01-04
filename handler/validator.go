package handler

import (
	"playbook-artifact-validator/ingress"
	probes "playbook-artifact-validator/instrumentation"
)

func validateArtifacts(request *ingress.Request) *ingress.Response {
	response := &ingress.Response{
		Request:    *request,
		Validation: "failure",
	}

	probes.ValidationFailed(request, "notImplemented")

	return response
}
