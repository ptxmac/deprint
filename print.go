package deprint

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Println format and print arguments like fmt.Println,
// but prefixes the output with file:line of the caller
func Println(a ...interface{}) (int, error) {
	if Output == Disabled {
		return 0, nil
	}
	str := fmt.Sprintln(a...)
	_, file, line, _ := runtime.Caller(1)
	_, file = filepath.Split(file)
	out := os.Stdout
	if Output == Stderr {
		out = os.Stderr
	}
	return fmt.Fprintf(out, "%s:%d: %s", file, line, str)
}
