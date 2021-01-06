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

func FetchArchiveError(request *ingress.Request, err error) {
	errorTotal.WithLabelValues("s3fetch").Inc()
	log.Errorw("Failed to fetch uploaded archive", "error", err, "reqId", request.RequestID)
}
