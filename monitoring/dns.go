package monitoring

import (
	"context"
	"log"
	"net"

	"github.com/mparvin/netinfo-exporter/config"
)

// PerformDNSLookup performs a DNS lookup for the given target and record type
func PerformDNSLookup(dnsConfig config.DNSConfig) bool {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), dnsConfig.Timeout)
	defer cancel()

	// Set up a custom resolver with the specified nameserver
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.DialTimeout(network, dnsConfig.Nameserver, dnsConfig.Timeout)
		},
	}

	// Perform the DNS lookup based on the record type
	var err error
	switch dnsConfig.RecordType {
	case "A":
		_, err = resolver.LookupHost(ctx, dnsConfig.Target)
	case "AAAA":
		_, err = resolver.LookupIPAddr(ctx, dnsConfig.Target)
	case "CNAME":
		_, err = resolver.LookupCNAME(ctx, dnsConfig.Target)
	case "MX":
		_, err = resolver.LookupMX(ctx, dnsConfig.Target)
	case "TXT":
		_, err = resolver.LookupTXT(ctx, dnsConfig.Target)
	default:
		log.Printf("Unsupported DNS record type: %s", dnsConfig.RecordType)
		return false
	}

	// Check if the lookup was successful
	if err != nil {
		log.Printf("DNS lookup failed for %s (type: %s, nameserver: %s): %v",
			dnsConfig.Target, dnsConfig.RecordType, dnsConfig.Nameserver, err)
		return false
	}

	log.Printf("DNS lookup succeeded for %s (type: %s, nameserver: %s)",
		dnsConfig.Target, dnsConfig.RecordType, dnsConfig.Nameserver)
	return true
}
