package fancyLogger

import (
	"fmt"
)

// At the core of logging all we want is something to log a message
// to something. We can achive this with fmt.println to the console.
// However you may want to log to something else like an ELK stack
// or structure your logs for processing later.
// This package provides a basic logger for you to work with but,
// also allows you to create your own functionality
// This is useful to retrofit logging functionality in a large codebase.
// Or to change logging behaviour in testing systems.

// DemoLogger - An interface that just describes a fancy loggers
// functions that are exposed.
type DemoLogger interface {
	Warn(string)
	Debug(string)
}

// FancyLogger is a struct that just holds your functions.
type FancyLogger struct {
	warnFunc  func(string)
	debugFunc func(string)
}

// Warn calls the your warn func and passes on the message.
func (l *FancyLogger) Warn(logMsg string) {
	l.warnFunc(logMsg)
}

// Debug calls your debug function and passes on the message.
func (l *FancyLogger) Debug(logMsg string) {
	l.debugFunc(logMsg)
}

// NewFancyLogger retruns a logger with your functions
func NewFancyLogger(warn, debug func(string)) *FancyLogger {
	return &FancyLogger{
		warnFunc:  warn,
		debugFunc: debug,
	}
}

// BoringLogger is just a plain logger
// Notice that we are using our NewFancyLogger constructor function
// and just putting in some basic logging methods.
func BoringLogger() *FancyLogger {
	return NewFancyLogger(
		func(msg string) {
			fmt.Printf("Warning: %s\n", msg)
		},
		func(msg string) {
			fmt.Printf("Debug: %s\n", msg)
		},
	)
}
