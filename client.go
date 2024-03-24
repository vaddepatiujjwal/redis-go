package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	// start the redis client
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer conn.Close()

	for i := 0; i < 1; i++ {
		fmt.Printf("made a call, %d\n", i)
		conn.Write([]byte("*2\r\n$4\r\necho\r\n$5\r\nheyrr\r\n"))
		time.Sleep(time.Second * 1)
	}
}
