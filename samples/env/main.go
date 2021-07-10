package main

import (
	"fmt"

	"go.ptx.dk/deprint"
)

func init() {
	deprint.FromEnv("DEBUG", deprint.Disabled)
}

func main() {
	fmt.Println("Only print debug statements if env DEBUG is set")
	fmt.Println("DEBUG=0 => print is disabled")
	fmt.Println("DEBUG=1 => print to stdout")
	fmt.Println("DEBUG=2 => print to stderr")

	deprint.Println("Debug was set to something!")
}
