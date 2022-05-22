package deprint

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func printWithCaller(str string) (int, error) {
	out := os.Stdout
	if Output == Stderr {
		out = os.Stderr
	}
	_, file, line, _ := runtime.Caller(2)
	_, file = filepath.Split(file)
	return fmt.Fprintf(out, "%s:%d: %s", file, line, str)
}

// Println format and print arguments like fmt.Println,
// but prefixes the output with file:line of the caller
func Println(a ...interface{}) (int, error) {
	if Output == Disabled {
		return 0, nil
	}
	str := fmt.Sprintln(a...)
	return printWithCaller(str)
}

// Printf formats and prints like fmt.Printf, but prefixes the output with file:line of the caller
func Printf(format string, a ...interface{}) (int, error) {
	if Output == Disabled {
		return 0, nil
	}
	str := fmt.Sprintf(format, a...)
	return printWithCaller(str)
}
