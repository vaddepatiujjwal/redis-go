package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var result = make(map[string]string)

func main() {

	// starts the redis server on it's default port
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer l.Close()

	for {
		// blocking call, this line waits untill we have a redis client makes a call
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		// Each client gets it's own go routine to serve concurrent redis client's
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// same client can make n numbers of command's
	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading command: ", err.Error())
		}

		input := string(buffer[:n])
		var response string

		if len(input) > 0 {
			//fmt.Printf("incoming command: %s\n", input)
			// extract command sent by client (e.g. Ping, Echo, Set or Get)
			tokens, _ := parseCommand(input)

			// handle redis commands
			switch strings.ToLower(tokens[0]) {
			case "echo":
				response = encodeRedisString(tokens[1])
			case "set":
				result[tokens[1]] = encodeRedisString(tokens[2])
				// result[tokens[3]] = tokens[4]
				response = "+OK\r\n"
			case "get":
				response = "$-1\r\n"

				if _, ok := result[tokens[1]]; ok {
					response = result[tokens[1]]
				}
			case "ping":
				response = "+PONG\r\n"
			}

			// write response to client,
			conn.Write([]byte(response))
		}
	}
}
