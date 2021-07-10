package main

import (
	"fmt"

	"go.ptx.dk/deprint"
)

func main() {
	fmt.Println("Normal print")
	deprint.Println("Debug print to stdout")
	deprint.Output = deprint.Disabled
	deprint.Println("Debug print to nowhere")
	deprint.Output = deprint.Stderr
	deprint.Println("Debug print to stderr")
}
