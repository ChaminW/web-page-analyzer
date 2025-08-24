package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	URLAnalysisTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "url_analysis_total",
		Help: "The total number of URL analyses performed",
	})

	URLAnalysisDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "url_analysis_duration_seconds",
		Help:    "Time taken to analyze URLs",
		Buckets: prometheus.DefBuckets,
	})

	URLAnalysisErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "url_analysis_errors_total",
		Help: "The total number of URL analysis errors",
	})
)
