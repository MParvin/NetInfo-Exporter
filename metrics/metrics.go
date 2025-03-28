package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	pingSuccess = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ping_success",
			Help: "Success status of ping",
		},
		[]string{"target"},
	)
	portCheckSuccess = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_check_success",
			Help: "Success status of port check",
		},
		[]string{"name"},
	)
	curlSuccess = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "curl_success",
			Help: "Success status of curl checks",
		},
		[]string{"url"},
	)
	dnsLookupSuccess = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "dns_lookup_success",
			Help: "Success status of DNS lookups",
		},
		[]string{"target"},
	)
)

func init() {
	prometheus.MustRegister(pingSuccess, portCheckSuccess, curlSuccess, dnsLookupSuccess)
}

func UpdatePingMetric(target string, success bool) {
	if success {
		pingSuccess.WithLabelValues(target).Set(1)
	} else {
		pingSuccess.WithLabelValues(target).Set(0)
	}
}

func UpdatePortCheckMetric(name string, success bool) {
	if success {
		portCheckSuccess.WithLabelValues(name).Set(1)
	} else {
		portCheckSuccess.WithLabelValues(name).Set(0)
	}
}

func UpdateCurlMetric(url string, success bool) {
	if success {
		curlSuccess.WithLabelValues(url).Set(1)
	} else {
		curlSuccess.WithLabelValues(url).Set(0)
	}
}

func UpdateDNSLookupMetric(target string, success bool) {
	if success {
		dnsLookupSuccess.WithLabelValues(target).Set(1)
	} else {
		dnsLookupSuccess.WithLabelValues(target).Set(0)
	}
}
