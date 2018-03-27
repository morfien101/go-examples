package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/morfien101/go-examples/unixSockets/client"
	"github.com/morfien101/go-examples/unixSockets/server"
)

var (
	socketLocation = flag.String("-s", "/tmp/app.sock", "Socket file location.")
)

// In this example we are trying to run 2 application code bases as a single unit.
// This is hacky but we get a pretty demo out of it. Lets work though the code.
func main() {
	// Collect the flags.
	flag.Parse()

	// We need to shutdown the services cleanly. *lies we actually just delete a file
	// but thats better than making the administrator or student do it.
	// This works by making use of the siganls package which is used to trap signals
	// from the OS. We then act upon them accordingly.
	// The way go does this is with channels.
	// So we make a channel to send signals down.
	c := make(chan os.Signal, 2)
	// We till the signals package where to send the signals. aka the channel we just made.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// We need a ticket later. But we also want to stop it when we clean up from a signal.
	// We create it here so it is in scope for the clean up function.
	ticker := time.NewTicker(time.Second * 3)

	// This our clean up function that does the work of shutting down our app.
	// It starts running from the very beginning but uses a go channels trick
	// to just sleep until it is told to work.
	go func() {
		// Recieving on an empty channel is a blocking action. However, we are in a
		// go routine so this is not really any issue to our application.
		// We will move past this once we receive something on this channel. The value
		// is just binned off. Think of it as a trip wire.
		<-c
		fmt.Println()
		fmt.Println("Caught a stop signal. Shutting down...")
		// Here is the ticket that we want to stop.
		// This will stop the client code from firing while we are trying to shutdown.
		ticker.Stop()
		// cleanup really just deletes the file we created for the socket.
		cleanup()
		// Exit happy once we are complete.
		os.Exit(0)
	}()

	// Start the server in a go routine so it acts as its own thing.
	go server.Main(*socketLocation)

	// Tickets are great. They stop you from writing sleep statements into your code.
	// They can also be controlled as we seen in the clean up section.
	// For a simple explanation this code is extreamly simular to this.
	// for {
	// 	  doSomething()
	// 	  time.Sleep(time.Second * 3)
	// }
	// With the added benifit that we can control the ticker.
	for _ = range ticker.C {
		client.Main(*socketLocation)
	}
}

// Delete the file so that it doesn't block the program from starting.
func cleanup() {
	err := os.Remove(*socketLocation)
	if err != nil {
		log.Println("Cloudn't clean up:", err)
	}
}
