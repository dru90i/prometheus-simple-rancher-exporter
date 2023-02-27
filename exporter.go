package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	gaugeVecs map[string]*prometheus.GaugeVec
}

func newExporter() *Exporter {
	gaugeVecs := addMetrics()
	return &Exporter{
		gaugeVecs: gaugeVecs,
	}
}
