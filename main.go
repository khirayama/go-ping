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

	ip, err := getIPAddr(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getIPAddr")
		os.Exit(1)
	}

	conn, err := net.Dial("ip4:1", ip.String()) // なんでip.String()してるんだ？
	if err != nil {
		fmt.Fprintf(os.Stderr, "net.Dial", err)
		os.Exit(1)
	}
	defer conn.Close() // main終了時に呼ばれる

	fmt.Println("PING", os.Args[1], "(", ip, ")")
}
