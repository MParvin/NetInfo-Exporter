package monitoring

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/mparvin/netinfo-exporter/metrics"
)

func CheckPort(name string, target string, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", target)
	if err != nil {
		log.Printf("Failed to connect to %s (%s): %v", name, target, err)
		metrics.UpdatePortCheckMetric(name, false)
		return false
	}
	defer conn.Close()

	log.Printf("Successfully connected to %s (%s)", name, target)
	metrics.UpdatePortCheckMetric(name, true)
	return true
}
