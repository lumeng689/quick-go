package metrics

import "github.com/prometheus/client_golang/prometheus"

var CpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

var CounterOps = prometheus.NewCounter(prometheus.CounterOpts{
	Subsystem: "sub_sys_1",
	Name:      "gin_counter_test1",
	Help:      "first counter test1",
})

func init() {
	prometheus.MustRegister(CpuTemp, CounterOps)
}
