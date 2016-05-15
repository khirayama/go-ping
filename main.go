package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
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

func pinger(conn net.Conn, id uint16, sigc chan os.Signal, c chan int) {
	nt := 0
	seq := uint16(0)
	t := time.NewTicker(1 * time.Second)
	done := false
	for !done {
		select {
		case <-sigc: // <- ってなんだったけ？
			done = true
		case <-t.C:
			tb, err := time.Now().MarshalBinary()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Time.MarshalBinary:", err)
				os.Exit(1)
			}
		}
	}
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

	// このブロックよく理解できてないな
	fmt.Println("PING", os.Args[1], "(", ip, ")")
	sigc := make(chan os.Signal, 1)   // makeはスライス、マップ、チャネルのみのnew的なもの
	signal.Notify(sigc, os.Interrupt) // https://golang.org/pkg/os/signal/#Notify
	c := make(chan int, 1)

	id := uint16(os.Getpid() & 0xfff)
	go pinger(conn, id, sigc, c)
}
