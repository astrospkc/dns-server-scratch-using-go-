package main

import (
	"dnsServer/cmd/dns"
	"fmt"

	// "sync"
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


	buf := make([]byte, 512)
	for{
		fmt.Print("hello")
		_, clientAddr ,_ :=conn.ReadFromUDP(buf)

		// parseHeader
		reqHeader , _:=dns.ParseDNSHeader(buf)
		q, _, _ := dns.ParseQuestion(buf, 12)
		fmt.Print("name: ", q.Name)
		resHeader := &dns.DNSHeader{
            ID: reqHeader.ID,
            Flags: 0x8180, // standard response
            QDCOUNT: 1,
            ANCOUNT: 1,
            NSCOUNT: 0,
            ARCOUNT: 0,
        }

		        // Build full response
        response := dns.BuildDNSHeader(resHeader)
        response = append(response, dns.WriteQuestion(q)...)

        // Lookup IP (for now: static)
        ip := "192.168.112.1"

        // Build answer section
        ans := dns.BuildARecord(12, ip, 60)
		fmt.Print("ans", ans)
        response = append(response, ans...)

        // Send response back
        conn.WriteToUDP(response, clientAddr)
	}

	

	

	
}


// var wg sync.WaitGroup

// func worker(id int) {
//     defer wg.Done()
//     fmt.Println("Worker", id)
// }

// func main() {
//     for i := 1; i <= 3; i++ {
//         wg.Add(1)
//         go worker(i)
//     }
//     wg.Wait() // wait for all workers
// }


