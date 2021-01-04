package instrumentation

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

var (
	validationSuccessTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "validation_success_total",
		Help: "The total number of successfully validated payloads",
	})

	validationFailureTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "validation_failure_total",
		Help: "The total number of payloads that did not pass validation",
	})

	errorTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "error_total",
		Help: "The total number of errors during payloads processing",
	}, []string{"phase"})
)

func StartMetrics(config *viper.Viper, mux *http.ServeMux) {
	mux.Handle(config.GetString("metrics.path"), promhttp.Handler())
}
