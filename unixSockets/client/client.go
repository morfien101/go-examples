package client

import (
	"fmt"
	"io"
	"log"
	"net"
)

// This is a simple client that will write to a unix socket and then wait for the reply.

// Main is the clients main function
func Main(socketLocation string) {
	// Create a connection to an existing socket.
	c, err := net.Dial("unix", socketLocation)
	if err != nil {
		log.Fatalln("Failed to connect to the socket:", err)
	}
	// Close the connection at the end.
	// We could just call the Close function. However I want to handle the error if
	// there is one. So annonymous function it is.
	defer func() {
		err := c.Close()
		if err != nil {
			fmt.Println("Client error:", err)
		}
	}()

	stringToWrite := []byte("GET /healthz")
	write(c, stringToWrite)
	read(c)
}

func write(c io.Writer, bs []byte) {
	// Write on the socket.
	_, err := c.Write(bs)
	if err != nil {
		log.Fatal("write error:", err)
	}
}

func read(r io.Reader) {
	// Make a buffer to read out the response. It will need to big enough to fit the reply.
	buf := make([]byte, 5120)
	n, err := r.Read(buf[:])
	// We need to check for an io.EOF here because the server closes the connection.
	// We expect this behavior.
	if err != nil && err != io.EOF {
		log.Fatalln("Failed to read from the socket:", err)
		return
	}
	// print out what we got from the server. This could be a function to digest it and
	// forward it onto a metric aggregation platform like influxdb or ELK.
	fmt.Println("Client got:", string(buf[0:n]))
}
