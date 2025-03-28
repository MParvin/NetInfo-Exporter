package monitoring

import (
	"log"
	"strings"
	"sync"

	"github.com/mparvin/netinfo-exporter/config"
	"github.com/mparvin/netinfo-exporter/metrics"
)

func RunAllChecks(cfg *config.Config) {
	var wg sync.WaitGroup

	for _, ping := range cfg.Ping {
		wg.Add(1)
		go func(ping config.PingConfig) {
			defer wg.Done()
			log.Printf("Running ping check for: %s with timeout: %s", ping.Target, ping.Timeout)
			success := PerformPing(ping.Target, ping.Timeout)
			metrics.UpdatePingMetric(ping.Target, success)
			log.Printf("Updated ping_success metric for target %s: %t", ping.Target, success)
		}(ping)
	}

	for _, port := range cfg.Port {
		wg.Add(1)
		go func(port config.PortConfig) {
			defer wg.Done()
			log.Printf("Running port check for: %s:%d with timeout: %s", port.Target, port.Port, port.Timeout)
			name := strings.ReplaceAll(strings.ReplaceAll(port.Target, ".", "_"), ":", "_")
			success := CheckPort(name, port.Target, port.Timeout)
			metrics.UpdatePortCheckMetric(name, success)
			log.Printf("Updated port_check_success metric for target %s:%d: %t", port.Target, port.Port, success)
		}(port)
	}

	for _, urlConfig := range cfg.URL {
		wg.Add(1)
		go func(urlConfig config.URLConfig) {
			defer wg.Done()
			log.Printf("Running HTTP check for: %s with method: %s, timeout: %s", urlConfig.Target, urlConfig.Method, urlConfig.Timeout)
			success := PerformHTTPCheck(urlConfig)
			metrics.UpdateCurlMetric(urlConfig.Target, success)
			log.Printf("Updated curl_success metric for target %s: %t", urlConfig.Target, success)
		}(urlConfig)
	}

	for _, dns := range cfg.DNS {
		wg.Add(1)
		go func(dns config.DNSConfig) {
			defer wg.Done()
			log.Printf("Running DNS check for: %s with record type: %s, nameserver: %s, timeout: %s",
				dns.Target, dns.RecordType, dns.Nameserver, dns.Timeout)
			success := PerformDNSLookup(dns)
			metrics.UpdateDNSLookupMetric(dns.Target, success)
			log.Printf("Updated dns_lookup_success metric for target %s: %t", dns.Target, success)
		}(dns)
	}

	wg.Wait()
}
