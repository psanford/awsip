package awsip

import (
	"fmt"
	"net/netip"
	"testing"
)

func TestIsAwsIP(t *testing.T) {
	awsIPs := []netip.Addr{
		netip.MustParseAddr("54.74.0.27"),
		netip.MustParseAddr("2a05:d03a:8000::1"),
	}

	for _, addr := range awsIPs {
		if !IsAwsIP(addr) {
			t.Errorf("Expected %s to match aws ip but did not", addr)
		}
	}

	nonAwsIPs := []netip.Addr{
		netip.MustParseAddr("127.0.0.12"),
		netip.MustParseAddr("10.48.20.96"),
		netip.MustParseAddr("8.8.8.8"),
		netip.MustParseAddr("2606:4700:4700::1111"),
	}
	for _, addr := range nonAwsIPs {
		if IsAwsIP(addr) {
			t.Errorf("%s is not an AWS ip address, but it matched", addr)
		}
	}
}

func ExampleRange() {
	ip := netip.MustParseAddr("54.74.0.27")
	r := Range(ip)
	fmt.Println(r.Prefix)
	fmt.Println(r.NetworkBorderGroup)
	fmt.Println(r.Region)
	fmt.Println(r.Services)
	// Output:
	// 54.74.0.0/15
	// eu-west-1
	// eu-west-1
	// [AMAZON EC2]
}

func ExampleIsAwsIP() {
	ips := []netip.Addr{
		netip.MustParseAddr("54.74.0.27"),
		netip.MustParseAddr("127.0.0.1"),
	}
	for _, ip := range ips {
		if IsAwsIP(ip) {
			fmt.Printf("%s is AWS\n", ip)
		} else {
			fmt.Printf("%s is NOT AWS\n", ip)
		}
	}
	// Output:
	// 54.74.0.27 is AWS
	// 127.0.0.1 is NOT AWS
}
