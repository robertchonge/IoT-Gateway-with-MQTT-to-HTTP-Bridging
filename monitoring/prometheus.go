package monitoring

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

func StartPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Prometheus metrics exposed on :9100/metrics")
	http.ListenAndServe(":9100", nil)
}
