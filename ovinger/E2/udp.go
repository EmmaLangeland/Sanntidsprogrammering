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
		Port: 20003,
		IP: net.IPv4zero //net.IPv4(10, 24, 35, 209), // IP: net.IPv4(10, 24, 35, 209), addr: 10.24.35.209:20004
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	buf:= make([]byte, 1024)
	for {
		n, fromWhoAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		localIP, err := localip.LocalIP()
		if (fromWhoAddr.IP != localIP){
			fmt.Printf("[UDP RECEIVER] From %s: %s\n", fromWhoAddr.String(), string(buf[:n]))
		}
	}
	
}

func UDP_Sender() {
	net.DialUDP("udp", ":20003")

}

func main() {
	go 

	select{}
}