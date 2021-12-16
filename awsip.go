package awsip

import (
	"net/netip"
)

//go:generate go run update_ips.go

// IsAwsIP returns true if the ip address falls within one of the known
// AWS ip ranges.
func IsAwsIP(ip netip.Addr) bool {
	r := Range(ip)
	return r != nil
}

// Range returns the ip range and metadata an address falls within.
// If the IP is not an AWS IP address it returns nil
func Range(ip netip.Addr) *IPRange {
	for _, r := range ipRanges {
		if r.Prefix.Contains(ip) {
			return &r
		}
	}
	return nil
}

type IPRange struct {
	Prefix             netip.Prefix
	NetworkBorderGroup string
	Region             string
	Service            string
}
