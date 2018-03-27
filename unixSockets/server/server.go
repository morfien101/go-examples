package server

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

// internalState is a container that we can store something in. In a real world case
// it could contain metric counters that you would expose to a monitoring platform.
// I have also put json labels on it for the marshaller later.
// It has a mutex as there could be multiple clients calling it or a write happening
// during a read. This basically just makes it thread safe.
type internalState struct {
	Ticks    int64 `json:"ticks"`
	LastTick int64 `json:"epoch_last_tick"`
	sync.RWMutex
}

// Easy function to just change the values. So we can see it changing in the demo.
func (is *internalState) addTick() {
	is.Lock()
	defer is.Unlock()
	is.Ticks++
	is.LastTick = time.Now().Unix()
}

// readState JSON Marshals our internalState and returns a byte slice with the data.
// If there is an error we return the error and an empty byte slice.
func (is *internalState) readState() ([]byte, error) {
	is.RLock()
	defer is.RUnlock()
	b, err := json.Marshal(is)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

// Loop and sleep for a random time. Increment the counters on each loop iteration.
// We are just simulating a task of some kind.
func (is *internalState) work(state *internalState) {
	for {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		is.addTick()
	}
}

// Main is the servers main function.
func Main(socketLocation string) {
	// Create state and make a worker do something to it.
	state := &internalState{}
	log.Println("Starting worker.")
	go state.work(state)

	// Now lets start up a unix socket based server.
	// Unix sockets are bidirectional. They are also faster than TCP as there is less
	// overhead to deal with. With Intel specter this could save lots of context switch
	// CPU overhead.
	log.Println("Starting lister on", socketLocation)
	l, err := net.Listen("unix", socketLocation)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// we need to listen for new connection requests.
	// Accept will block until a request comes in, this means this for loop will just wait
	// until there is a connection to handle.
	for {
		clientConnection, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		// Kick off a go routine to handle the connection. Then go back to waiting for the next.
		log.Println("Starting a handler for the connection.")
		go echoServer(clientConnection, state)
	}
}

func echoServer(c net.Conn, state *internalState) {
	// Close the connection after the request is finished.
	// Sockets are bidirectional and persistant. We could have the server and client
	// keep talking over the same connection. However in this case I am closing the
	// socket to signal that I am done replying. Since we are using this as a health
	// checking service we just reply with the health and then close the connection.
	defer c.Close()

	// We need a buffer to store the reqeust in as we pull it off the connection.
	buf := make([]byte, 512)
	// Read will copy the bytes to the buffer until it finds a EOF.
	// It will return a counter to let to know how many bytes you have in the buffer.
	// In this example the buffer is a fixed size. This isn't an issue if you are sure
	// of what is going to come into the socket.
	bytesRead, err := c.Read(buf)
	if err != nil {
		log.Println("Error reading from the socket. Error:", err)
		return
	}

	// We drop the last character if it is going to be a newline char
	if buf[bytesRead-1 : bytesRead][0] == byte('\n') {
		bytesRead--
	}
	// Side note here; If this was a long running connection and the buffer was huge,
	// we would want ot clean out the buf var as it is held in memory for the duration
	// of this function.
	data := buf[0:bytesRead]
	log.Println("Server got:", string(data))
	// Is the string what we expect?
	if string(data) == "GET /healthz" {
		stateData, err := state.readState()
		if err != nil {
			log.Println("State Marshal error:", err)
			return
		}
		// I want to send a new line to the client with our response. So I just stick
		// it onto the end of the slice. Make it nice to look at in the console. In
		// the real world you may not want to do this.
		stateData = append(stateData, '\n')
		// Lucky for us... The c.Write is a Writer interface that you just give a []byte
		// and it consumes it. It will return an error if something goes wrong and how
		// many bytes are written, you could use this to error check. eg. Did I send as
		// many bytes as there is in the slice I was sending?
		writtenBytes, err := c.Write(stateData)
		if err != nil {
			log.Println("Write: ", err)
			return
		}

		// Some logging to see in the console.
		log.Printf("Server responded with %d bytes.\n", writtenBytes)
	} else {
		log.Println(string(data), "is not valid.")
	}
	// At this point the function is finished and the defer statement will kick in to
	// close the connection.
}
