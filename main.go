package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func getIPAddr(host string) (net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.To4(), nil
		}
	}

	return nil, errors.New("IP address not found")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "arg error")
		os.Exit(1)
	}
	host := os.Args[1]

	// ip, err := getIPAddr(host)
	ip, _ := getIPAddr(host)
	fmt.Println(ip)
}
