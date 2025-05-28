package main

import (
	"Network-go/network/localip"
	"fmt"
	"net"
	"os"
	"time"
)

func UDP_Receiver() {

	addr := net.UDPAddr{
		IP:   net.IPv4zero, //net.IPv4(10, 24, 35, 209), // IP: net.IPv4(10, 24, 35, 209), addr: 10.24.35.209:20004
		Port: 20000,
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		numBytesReceived, fromWhoAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		localIP, err := localip.LocalIP()
		if fromWhoAddr.IP.String() != localIP {
			fmt.Printf("[UDP RECEIVER] From %s: %s\n", fromWhoAddr.String(), string(buf[:numBytesReceived]))
		}
	}
}

func UDP_Sender() {
	senderAddr := net.UDPAddr{
		IP: net.IPv4(10, 24, 32, 55) //net.IPv4(10, 24, 35, 209)
		Port: 20001
	}

	conn, err := net.DialUDP("udp", nil, &senderAddr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	message := "DU STÅR EKSAMEN \x00"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending:", err)
		return
	} else {
		fmt.Printf("[UDP SENDER] Sent to %s:%d: %s\n", senderAddr.IP.String(), message)
	}

}

func main() {
	go UDP_Receiver()

	time.Sleep(1 * time.Second)

	UDP_Sender()

	select {}
}
