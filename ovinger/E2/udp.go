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
		Port: 20000,
		IP:   net.IPv4zero, //net.IPv4(10, 24, 35, 209), // IP: net.IPv4(10, 24, 35, 209), addr: 10.24.35.209:20004
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
		senderIP: net.IPv4(10, 24, 35, 255) //net.IPv4(10, 24, 35, 209)
		senderPort: 20004
	}

	conn, err := net.DialUDP("udp", nil, &senderAddr)
	defer conn.Close()

	message := "DU STÅR EKSAMEN \x00"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending:", err)
		return
	} else {
		fmt.Printf("[UDP SENDER] Sent to %s:%d: %s\n", serverIP.String(), serverPort, message)
	}

}

func main() {
	go UDP_Receiver()

	time.Sleep(1 * time.Second)

	UDP_Sender()

	select {}
}
