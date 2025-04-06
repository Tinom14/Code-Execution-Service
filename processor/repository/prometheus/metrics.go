package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type PrometheusStorage struct {
	taskDuration  *prometheus.HistogramVec
	languageUsage *prometheus.CounterVec
}

func NewPrometheusStorage() *PrometheusStorage {
	return &PrometheusStorage{
		taskDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "task_duration_seconds",
				Help: "Time taken to process a task",
			},
			[]string{"language"},
		),
		languageUsage: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "language_usage_total",
				Help: "Count of tasks by language",
			},
			[]string{"language"},
		),
	}
}

func (m *PrometheusStorage) Register() {
	prometheus.MustRegister(m.taskDuration)
	prometheus.MustRegister(m.languageUsage)
}

func (m *PrometheusStorage) RecordTaskDuration(language string, duration time.Duration) {
	m.taskDuration.WithLabelValues(language).Observe(duration.Seconds())
}

func (m *PrometheusStorage) RecordLanguageUsage(language string) {
	m.languageUsage.WithLabelValues(language).Inc()
}
