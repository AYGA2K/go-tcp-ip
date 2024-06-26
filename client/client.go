package main

import (
	"fmt"
	"net"
)

func main() {
	client()
}

func client() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and process data from the server
	// ...
}
