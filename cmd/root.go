/*
Copyright © 2024 NAME HERE mparvin@parsops.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mparvin/netinfo-exporter/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile           string
	listenAddress     string
	defaultListenAddr = ":9876"
)

var rootCmd = &cobra.Command{
	Use:   "netinfo_exporter",
	Short: "NetInfo Exporter is a Prometheus exporter for monitoring network health and performance",
	Long: `NetInfo Exporter performs various network checks including ping, port availability,
URL health, and DNS lookups. It provides Prometheus metrics for monitoring.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			log.Fatal("Configuration file is required. Use --config to specify the path.")
		}

		cfg, err := config.LoadConfig(cfgFile) // Load the configuration
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}

		fmt.Printf("Starting NetInfo Exporter on %s with config: %s\n", cfg.ListenAddress, cfgFile)

		// Start the Prometheus metrics exporter
		startExporter(*cfg)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to the configuration file (required)")
	rootCmd.PersistentFlags().StringVar(&listenAddress, "web.listen-address", defaultListenAddr, "Address to listen on for Prometheus metrics")

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("web.listen-address", rootCmd.PersistentFlags().Lookup("web.listen-address"))
}
