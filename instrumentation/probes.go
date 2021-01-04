package instrumentation

import (
	"playbook-artifact-validator/ingress"
	"playbook-artifact-validator/utils"
)

var log = utils.GetLoggerOrDie()

func ValidationFailed(request *ingress.Request, cause string) {
	validationFailureTotal.Inc()
	log.Infow("Payload validation failed", "cause", cause, "reqId", request.RequestID)
}

func UnmarshallingError(err error) {
	errorTotal.WithLabelValues("unmarshall").Inc()
	log.Errorw("Request unmarshalling failed", "error", err) // TODO some correlation info
}
