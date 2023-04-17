package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func GetHistogramValue(col prometheus.Histogram) float64 {
	c := make(chan prometheus.Metric, 1)
	col.Collect(c)
	m := dto.Metric{}
	_ = (<-c).Write(&m)
	return float64(m.GetHistogram().GetSampleCount())
}

func GetCounterValue(col prometheus.Counter) float64 {
	c := make(chan prometheus.Metric, 1)
	col.Collect(c)
	m := dto.Metric{}
	_ = (<-c).Write(&m)
	return m.GetCounter().GetValue()
}

func GetGaugeValue(col prometheus.Gauge) float64 {
	c := make(chan prometheus.Metric, 1)
	col.Collect(c)
	m := dto.Metric{}
	_ = (<-c).Write(&m)
	return m.GetGauge().GetValue()
}
