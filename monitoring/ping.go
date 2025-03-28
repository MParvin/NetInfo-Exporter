package monitoring

import (
	"log"
	"time"

	"github.com/go-ping/ping"

	"github.com/mparvin/netinfo-exporter/metrics"
)

func PerformPing(target string, timeout time.Duration) bool {
	log.Printf("Pinging target: %s with timeout: %v", target, timeout)

	pinger, err := ping.NewPinger(target)
	if err != nil {
		log.Printf("Failed to create pinger for target %s: %v", target, err)
		metrics.UpdatePingMetric(target, false)
		return false
	}

	pinger.Count = 3
	pinger.Timeout = timeout

	err = pinger.Run()
	if err != nil {
		log.Printf("Ping failed for target %s: %v", target, err)
		metrics.UpdatePingMetric(target, false)
		return false
	}

	stats := pinger.Statistics()
	success := stats.PacketsRecv > 0

	metrics.UpdatePingMetric(target, success)

	return success
}
