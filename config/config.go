package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type PingConfig struct {
	Target  string        `mapstructure:"target"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type PortConfig struct {
	Target  string        `mapstructure:"target"`
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type URLConfig struct {
	Target         string        `mapstructure:"target"`
	Method         string        `mapstructure:"method"`
	Timeout        time.Duration `mapstructure:"timeout"`
	ExpectedStatus int           `mapstructure:"expected_status"`
	VerifySSL      bool          `mapstructure:"verify_ssl"`
}

type DNSConfig struct {
	Target     string        `mapstructure:"target"`
	RecordType string        `mapstructure:"record_type"`
	Nameserver string        `mapstructure:"nameserver"`
	Timeout    time.Duration `mapstructure:"timeout"`
}

type Config struct {
	ListenAddress string       `mapstructure:"listen_address"`
	Ping          []PingConfig `mapstructure:"ping"`
	Port          []PortConfig `mapstructure:"port"`
	URL           []URLConfig  `mapstructure:"url"`
	DNS           []DNSConfig  `mapstructure:"dns"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}
