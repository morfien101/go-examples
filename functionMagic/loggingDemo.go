package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/morfien101/go-examples/functionMagic/fancyLogger"
)

// loggingDemo is a function that shows a more realistic way of using
// functions that return a function.
func loggingDemo() {
	// Both the loggers below impliment the methods to statisfy the
	// fancyLogger.DemoLogger interface. The kicker is that package
	// fancyLogger actually makes these for us and sends them to us.
	// So it satisfies itself...

	// We need a toggle to turn on and off debug logging when we want
	debugToggle := true
	// This is our fancy logger. It will log in JSON so it is easy to
	// digest in something like ELK. It also makes use of a toggle for
	// debug messages and tracks how many messages it has logged.
	// Fancy right!!!
	jlgr := newJSONLogger(&debugToggle, 0)

	// BoringLogger is just a very basic logger. It just prints messages.
	flgr := fancyLogger.BoringLogger()

	// Lets see them in action. You will need to run this to see the output
	fmt.Printf("\n\nLoggers - Debug1\n")
	var m string
	m = "Oh No! It's all going wrong!"
	jlgr.Debug(m)
	flgr.Debug(m)

	fmt.Printf("\n\nLoggers - Debug2\n")
	m = "Hide the evidence!!!"
	debugToggle = false
	jlgr.Debug(m)
	flgr.Debug(m)

	fmt.Printf("\n\nLoggers - Warning\n")
	m = "Things could be going better"
	jlgr.Warn(m)
	flgr.Warn(m)
}

// The fancyLogger package doesn't know what debugEnabled and logIDTracker are
// so it wouldn't know what to do with it. This is called dependency injection.
// We are giving the logger something to work with that the original description
// didn't specify for. This is not a bad thing. It is actually a great thing. It
// affords us the freedom to do things like the below.
func newJSONLogger(debugEnabled *bool, logIDTracker int64) fancyLogger.DemoLogger {
	// fancyLogger.NewFancyLogger just returns a struct that has 2 functions. Warn()
	// Debug(). However it takes 2 functions and allocates them to these functions.
	return fancyLogger.NewFancyLogger(
		// This is our Warn function.
		func(msg string) {
			// Pull the current ID and then increment it.
			nextID := logIDTracker
			logIDTracker++
			// Make a struct and lable it for the JSON Marshaller
			log := &struct {
				LogType  string `json:"log_type"`
				LogID    int64  `json:"log_id"`
				LogEpoch int64  `json:"log_epoch"`
				LogMsg   string `json:"log_message"`
			}{
				// Set the values to what you want.
				LogType:  "WARN",
				LogID:    nextID,
				LogEpoch: time.Now().Unix(),
				LogMsg:   msg,
			}
			// Send it to be marshaled into JSON. We wget back an error from the Marshaller
			// and a []byte containing our data.
			output, err := json.Marshal(log)
			if err != nil {
				// To be fair we handle this pretty poorly as we just print to the console
				// and then return. But its a demo so thats ok.
				fmt.Printf("Failed to create log. Error: %s", err)
				return
			}
			// Now print to the console our wonderful log message in JSON.
			// This could also be a delivery method to a logging service.
			// HTTP Post or message on a AMQ queue?
			fmt.Println(string(output))
		},
		// This is the debug function.
		func(msg string) {
			// We only go into this if we need. Check if the pointer of the toggle
			// is true.
			if *debugEnabled {
				nextID := logIDTracker
				logIDTracker++
				// Notice that we have more fields on this log. It's debug so get everything.
				log := &struct {
					LogType       string `json:"log_type"`
					LogID         int64  `json:"log_id"`
					LogEpoch      int64  `json:"log_epoch"`
					LogMsg        string `json:"log_message"`
					LogGoVersion  string `json:"go_version"`
					LogGoRoutines int    `json:"go_number_of_goroutines"`
					LogGOOS       string `json:"os"`
				}{
					LogType:  "Debug",
					LogID:    nextID,
					LogEpoch: time.Now().Unix(),
					LogMsg:   msg,
					// The runtime package is good for gathering details of whats happening
					// inside the go engine.
					LogGoVersion:  runtime.Version(),
					LogGoRoutines: runtime.NumGoroutine(),
					LogGOOS:       runtime.GOOS,
				}
				// The rest is the same as above.
				output, err := json.Marshal(log)
				if err != nil {
					fmt.Printf("Failed to create log. Error: %s", err)
					return
				}
				fmt.Println(string(output))
			}
		},
	)
}
