package monitoring

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/mparvin/netinfo-exporter/config"
	"github.com/mparvin/netinfo-exporter/metrics"
)

func PerformHTTPCheck(urlConfig config.URLConfig) bool {
	// Log the HTTP check
	log.Printf("Performing HTTP check (PerformHTTPCheck) for: %s with method: %s", urlConfig.Target, urlConfig.Method)

	// Use timeout directly from the configuration
	timeout := urlConfig.Timeout

	// Create an HTTP client with a timeout and optional SSL verification
	client := &http.Client{
		Timeout: timeout,
	}
	if !urlConfig.VerifySSL {
		// Disable SSL verification if specified in the config
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	// Create the HTTP request
	req, err := http.NewRequest(urlConfig.Method, urlConfig.Target, nil)
	if err != nil {
		log.Printf("Failed to create HTTP request for %s: %v", urlConfig.Target, err)
		metrics.UpdateCurlMetric(urlConfig.Target, false)
		return false
	}

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("HTTP check failed for %s: %v", urlConfig.Target, err)
		metrics.UpdateCurlMetric(urlConfig.Target, false)
		return false
	}
	defer resp.Body.Close()

	// Check if the status code matches the expected status
	success := resp.StatusCode == urlConfig.ExpectedStatus
	if success {
		log.Printf("HTTP check succeeded for %s with status code: %d", urlConfig.Target, resp.StatusCode)
	} else {
		log.Printf("HTTP check failed for %s with status code: %d (expected: %d)", urlConfig.Target, resp.StatusCode, urlConfig.ExpectedStatus)
	}

	// Update the curl metric
	metrics.UpdateCurlMetric(urlConfig.Target, success)

	return success
}
