package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// handleConnection processes a single TCP connection
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading from connection: %v", err)
			return
		}

		fmt.Printf("Message Received from %s: %s", conn.RemoteAddr(), string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}

// ServerStart starts the TCP server in a new goroutine
func ServerStart() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Printf("Error starting TCP server: %v", err)
		return
	}
	defer ln.Close()

	fmt.Println("TCP Server listening on port 8000")

	// Accept connections in a loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle each connection in a new goroutine
		go handleConnection(conn)
	}
}

// ClientStart starts the TCP client
func ClientStart() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			return
		}

		// Send message to server
		_, err = fmt.Fprintf(conn, text)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}

		// Read response from server
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Error reading response: %v", err)
			return
		}

		fmt.Print("Message from server: " + message)
	}
}
