package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/netip"
	"os"

	"github.com/psanford/awsip"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <ip>\n", os.Args[0])
	}

	ipStr := os.Args[1]

	addr, err := netip.ParseAddr(ipStr)
	if err != nil {
		log.Fatalf("ip parse error: %s", err)
	}

	r := awsip.Range(addr)
	if r == nil {
		os.Exit(1)
	}

	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
