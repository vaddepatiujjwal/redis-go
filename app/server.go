package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// start the redis server
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		// input := string(buffer[:n])

		fmt.Println("incoming command: ping")

		if err != nil {
			fmt.Println("Error reading command: ", err.Error())
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}
