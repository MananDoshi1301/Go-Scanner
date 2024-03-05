package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"time"	
)

func main() {
	var host string
	var port int

	// Default host: localhost
	flag.StringVar(&host, "host", "127.0.0.1", "Host for scanning ports")

	// Default port for sending tcp handshake
	flag.IntVar(&port, "port", 3306, "Port for scanning MySql")
	flag.Parse()


	log.Println("Search for MySql Begins...")
	defer log.Println("Search for MySql Ends...")

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		log.Printf("Failed to connect to %s: %v\n", address, err)
		return
	}

	defer conn.Close()

	// Attempt to initiate MySQL handshake	
	_, err = conn.Write([]byte{0x0a})	
	if err != nil {
		log.Fatal("Failed to write handshake packet:", err)
	}

	// Read the initial response 
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		log.Printf("Failed to read initial response packet: %v\n", err)
		return
	}

	// Extract server version
	versionStart := bytes.IndexByte(buffer, 10) + 1
	versionEnd := bytes.IndexByte(buffer[versionStart:], 0) + versionStart
	version := string(buffer[versionStart:versionEnd])

	// Final logs
	log.Println("MySQL appears to be running on the specified host and port:", address)
	log.Println("MySQL Server Version:", version)	
	log.Println("Complete Response:", string(buffer))	
}
