# NetInfo Exporter

NetInfo Exporter is a Prometheus exporter for monitoring network health and performance. It performs various network checks including ping, port availability, URL health, and DNS lookups.

## Features

- Configuration-driven checks using YAML
- Multiple check types supported:
  - ICMP Ping
  - TCP Port checks
  - URL health checks
  - DNS lookups
- Prometheus metrics for check status, duration, and errors
- Configurable check intervals and timeouts
- Comprehensive error handling and logging
- Modular and extensible design

## Installation

```bash
go install github.com/MParvin/NetInfo-Exporter
```

## Usage

1. Create a configuration file (see `config-example.yml` for example)
2. Set the `config` parameter to the path of the configuration file
2. Run the exporter:

```bash
netinfo_exporter --config=/path/to/config.yml --web.listen-address=:9876
```

## Configuration example


The configuration file uses YAML format and supports the following check types:

### Ping Checks
```yaml
checks:
  ping:
    - target: "8.8.8.8"
      timeout: 5s
```

### Port Checks
```yaml
checks:
  port:
    - target: "example.com"
      port: 443
      timeout: 5s
```

### URL Checks
```yaml
checks:
  url:
    - target: "https://example.com"
      method: "GET"
      timeout: 10s
      expected_status: 200
      verify_ssl: true
```

### DNS Checks
```yaml
checks:
  dns:
    - target: "example.com"
      record_type: "A"
      nameserver: "8.8.8.8:53"
      timeout: 5s
```

## Metrics

The exporter provides the following metrics:

- `netinfo_check_status`: Status of the check (1 for success, 0 for failure)
- `netinfo_check_duration_seconds`: Duration of the check in seconds
- `netinfo_check_errors_total`: Total number of check errors

All metrics include labels for `check_name` and `check_type`. Additional labels may be present depending on the check type.

## Contributing
Are you interested in contributing to this project?
Feel free to fork the repository and submit a pull request! ;)


## License

This project is licensed under the MIT License - see the LICENSE file for details.