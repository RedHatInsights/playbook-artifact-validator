package instrumentation

import (
	"playbook-artifact-validator/ingress"
	"playbook-artifact-validator/utils"
)

var log = utils.GetLoggerOrDie()

func ValidationSuccess(request *ingress.Request) {
	validationSuccessTotal.Inc()
	log.Debugw("Payload valid", "reqId", request.RequestID)
}

func ValidationFailed(request *ingress.Request, cause error) {
	validationFailureTotal.Inc()
	log.Infow("Rejecting payload due to validation failure", "cause", cause, "reqId", request.RequestID)
}

func UnmarshallingError(err error) {
	errorTotal.WithLabelValues("unmarshall").Inc()
	log.Errorw("Message unmarshalling failed", "error", err) // TODO some correlation info
}

func FetchArchiveError(request *ingress.Request, err error) {
	errorTotal.WithLabelValues("s3fetch").Inc()
	log.Errorw("Failed to fetch uploaded archive", "error", err, "reqId", request.RequestID)
}
