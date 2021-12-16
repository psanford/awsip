# awsip: a Go package to check if an IP belongs to AWS

awsip is a Go package that allows you to determine if an IP address belongs to AWS.

A cli tool is also included in `cmd/awsip` for easily checking the status of an ip address.

## Example:

```
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

func ExampleRange() {
	ip := netip.MustParseAddr("54.74.0.27")
	r := Range(ip)
	fmt.Println(r.Prefix)
	fmt.Println(r.NetworkBorderGroup)
	fmt.Println(r.Region)
	fmt.Println(r.Service)
	// Output:
	// 54.74.0.0/15
	// eu-west-1
	// eu-west-1
	// AMAZON
}
```

CLI:
```
$ ./awsip 54.74.0.27
{
  "Prefix": "54.74.0.0/15",
  "NetworkBorderGroup": "eu-west-1",
  "Region": "eu-west-1",
  "Service": "EC2"
}
```

## Updating the ip ranges

To update the ip ranges run: `go generate`. This will fetch from https://ip-ranges.amazonaws.com/ip-ranges.json and regenerate the `ips.gen.go` file.

## License

MIT
