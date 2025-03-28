/*
Copyright © 2024 NAME HERE mparvin@parsops.com
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/mparvin/netinfo-exporter/config"
	"github.com/mparvin/netinfo-exporter/monitoring"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func startExporter(cfg config.Config) {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Running all checks before serving metrics...")
		monitoring.RunAllChecks(&cfg)
		log.Println("All checks completed. Serving metrics...")
		promhttp.Handler().ServeHTTP(w, r)
	})

	log.Printf("Starting HTTP server on %s", cfg.ListenAddress)
	err := http.ListenAndServe(cfg.ListenAddress, nil)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
