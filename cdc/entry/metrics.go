// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package entry

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	mountDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "ticdc",
			Subsystem: "mounter",
			Name:      "unmarshal_and_mount",
			Help:      "Bucketed histogram of processing time (s) of unmarshal and mount in mounter.",
			Buckets:   prometheus.ExponentialBuckets(0.000001, 10, 10),
		}, []string{"namespace", "changefeed"})
	totalRowsCountGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "ticdc",
			Subsystem: "mounter",
			Name:      "total_rows_count",
			Help:      "The total count of rows that are processed by mounter",
		}, []string{"namespace", "changefeed"})
	ignoredDMLEventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "ticdc",
			Subsystem: "mounter",
			Name:      "ignored_dml_event_count",
			Help:      "The total count of dml events that are ignored in mounter.",
		}, []string{"namespace", "changefeed"})
)

// InitMetrics registers all metrics in this file
func InitMetrics(registry *prometheus.Registry) {
	registry.MustRegister(mountDuration)
	registry.MustRegister(totalRowsCountGauge)
	registry.MustRegister(ignoredDMLEventCounter)
}
