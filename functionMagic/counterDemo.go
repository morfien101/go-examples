package main

import "fmt"

// CounterDemo is just a holing function to make the code easier to read.
func counterDemo() {
	// Here we have 2 functions that return a function and allocates them
	// to a variable. The function funcdimentally is just a pointer to
	// some code. The () signals the compiler that we want to execute the
	// code in the pointer address.
	//
	// Running this code will produce very different results for both the
	// returned functions. Read further down to see why.
	f1 := newStaticCounter()
	f2 := newDynamicCounter()
	for a := 0; a < 10; a++ {
		f1()
		f2()
	}
}

// In this function we return the address of a function that we have already
// created. So it is a static function. It's named and everything.
func newStaticCounter() func() {
	return counter
}

// Counter just prints to the console.
// It has no choice but to redeclare the number every time or put
// the number into the global scope where everything can use it.
func counter() {
	i := 0
	fmt.Printf("Static: %d\n", i)
}

// In this function we return an anonymous function. It is scoped inside
// this function block. "i" is defined outside of the anonymous function,
// but is still only in scope in the newDynamicCounter function. Therefore,
// it is still considered private. Inside the anonymous function we refference
// "i" again. This linking means that the garbage collector will leave it alone.
// It also means that we read the same variable all the time allowing us
// to increment it on every call.
// The main counterDemo function knows nothing about this variable.
func newDynamicCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Printf("Dynamic: %d\n", i)
	}
}
