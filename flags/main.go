package main

import (
	"flag"
	"fmt"
)

var (
	version     = "0.0.1"
	versionFlag = flag.Bool("v", false, "Shows the version of the program")
	wordFlag    = flag.String("word", "", "word to echo back.")
)

func main() {
	flag.Parse()

	if *versionFlag {
		printVersion()
	} else {
		fmt.Println("No version flag given")
	}

	if len(*wordFlag) > 0 {
		fmt.Println("You passed in:", *wordFlag)
	} else {
		fmt.Println("No word passed in.")
	}

}

func printVersion() {
	fmt.Printf("The version is %s\n", version)
}
