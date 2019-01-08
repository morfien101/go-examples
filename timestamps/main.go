package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now().Format(time.RFC850)
  fmt.Println(t)
}
