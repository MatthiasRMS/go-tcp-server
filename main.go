package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Process incoming request in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Set a 10 seconds timeout
	timeoutDelay := 10 * time.Second
	conn.SetReadDeadline(time.Now().Add(timeoutDelay))
	defer conn.Close()

	// Write to connection
	fmt.Fprintln(conn, "Connection accepted !")

	// Read incoming messages
	bufferScanner := bufio.NewScanner(conn)
	for bufferScanner.Scan() {
		messageReceived := bufferScanner.Text()
		fmt.Println(messageReceived)
	}

	fmt.Println("Timed out!")
}
