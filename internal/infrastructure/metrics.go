package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	URLAnalysisTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_analysis",
		Help: "The total number of URL analyses performed",
	})

	URLAnalysisDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "analysis_duration_seconds",
		Help:    "Time taken to analyze URLs",
		Buckets: prometheus.DefBuckets,
	})

	URLAnalysisErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "analysis_errors_total",
		Help: "The total number of URL analysis errors",
	})
)
