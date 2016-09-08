// This program is used to print out the
// Fibanacci sequence starting from 1 and 1.
// The numbers get rather large and start looking
// a bit strange after 92. It is not this
// programs job to look into that. Rather
// just print out the numbers.
// You will need to pass in a positive
// whole number as an argument to get it to 
// print that many numbers.
package main

import(
	"fmt"		// Used to print to console.
	"strconv"	// Used to convert the string based arguments to ints
	"os"		// Used to read the arguments passed in.
)

func main() {
	// Check that a argument was passed in.
	if len(os.Args) < 2 {
		panic("You have not supplied a number to count to.")
	}
	// convert and check for errors. Errors mean a number was not 
	// passed in.
	iterations,err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("You have not passed a whole number in.")
	}
	// Warn of the strange numbers.
	if iterations > 92 {
		fmt.Println("You will get some strange looking numbers, but you asked for them.\n\n")
	}
	// Bucket to hold the numbers that we are currently working with.
	// Preseeded with our numbers.
	last_numbers := []int64{1,1}
	// Loop as many times as requested and print the numbers.
	for i := 1; i <= iterations; i++ {
		// Print the number first
		fmt.Printf("%d",last_numbers[0])
		// Print a , unless it is the last number in the list.
		if i != iterations{ fmt.Printf(", ") }
		// Update the vaules then go again.
		last_numbers[0], last_numbers[1] = last_numbers[1], last_numbers[0] + last_numbers[1]
	}
}