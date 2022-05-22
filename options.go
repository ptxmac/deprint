package deprint

import "os"

// OutputType selects where to output print statements
type OutputType int

const (
	// Stdout represents os.Stdout
	Stdout OutputType = iota
	// Stderr represents os.Stderr
	Stderr
	// Disabled disables all printing
	Disabled
)

// Output controls if and where the statements are printed
var Output = Stdout

const (
	EnvValueDisable = "0"
	EnvValueStdout  = "1"
	EnvValueStderr  = "2"
)

// FromEnv configures Output based on the environment variable named envName.
// envName=0 => disable, 1 => stdout, 2 => stderr
// defaults to fallback if the variable is not set.
func FromEnv(envName string, fallback OutputType) {
	switch os.Getenv(envName) {
	case EnvValueDisable:
		Output = Disabled
	case EnvValueStdout:
		Output = Stdout
	case EnvValueStderr:
		Output = Stderr
	default:
		Output = fallback
	}
}
