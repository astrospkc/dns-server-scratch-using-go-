package main

import (
	"fmt"
	"log"
	"net"
)

const (
	serverAddr = ":8080"
)

func main() {
	
	udpAddr , err := net.ResolveUDPAddr("udp",serverAddr)
	if err!=nil{
		log.Fatal("Failed to resolve udp addr", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err!=nil{
		log.Fatal("Failed to establised udp connection", err)
	}

	defer conn.Close()
	fmt.Printf("Udp server  listenig on %v\n", conn.LocalAddr().String())
}
