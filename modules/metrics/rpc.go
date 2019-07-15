package metrics

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/juju/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"reflect"
	"safebox.jerson.dev/api/modules/context"
	"time"
)

var buckets = []float64{50, 100, 300, 500, 1000, 3000, 5000}

var httpReqs = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
	},
	[]string{"code", "method", "path"},
)
var httpLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_milliseconds",
		Help:    "How long it took to process the request, partitioned by status code, method and HTTP path.",
		Buckets: buckets,
	},
	[]string{"code", "method", "path"},
)

// RPC ...
func RPC(ctx context.Base, service *rpc.HTTPService) {

	prometheus.MustRegister(httpReqs)
	prometheus.MustRegister(httpLatency)

	service.Use(func(name string, args []reflect.Value, context rpc.Context, next rpc.NextInvokeHandler) (results []reflect.Value, err error) {
		log := ctx.GetLogger(name)
		log.Debug("start")
		from := time.Now()
		status := http.StatusOK

		results, err = next(name, args, context)
		if err != nil {
			log.Error(err)

			if errors.IsNotFound(err) {
				status = http.StatusNotFound
			} else if errors.IsBadRequest(err) {
				status = http.StatusBadRequest
			} else if errors.IsNotImplemented(err) {
				status = http.StatusNotImplemented
			} else if errors.IsUnauthorized(err) {
				status = http.StatusUnauthorized
			} else {
				status = http.StatusInternalServerError
			}
		}

		timeSince := float64(time.Now().Sub(from).Nanoseconds()) / 1000000.0
		httpReqs.WithLabelValues(fmt.Sprint(status), http.MethodGet, name).Inc()
		httpLatency.WithLabelValues(fmt.Sprint(status), http.MethodGet, name).Observe(timeSince)

		return results, err
	})

	http.Handle("/metrics", promhttp.Handler())
}
