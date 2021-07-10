# deprint

[![Go Reference](https://pkg.go.dev/badge/go.ptx.dk/deprint.svg)](https://pkg.go.dev/go.ptx.dk/deprint)
[![Github Workflow](https://github.com/ptxmac/deprint/actions/workflows/go.yml/badge.svg)](https://github.com/ptxmac/deprint/actions/workflows/go.yml)

A small no-dependency library for managing debug print statements in golang projects

## Sample

```go
package sample

import "go.ptx.dk/deprint"

func ExamplePrintln() {
	deprint.Println("Hello world")
	// Output: print_test.go:10: Hello world
}

```

## Usage

`go get -u go.ptx.dk/deprint`

Use `deprint.Println` anywhere you would normally use `fmt.Println`.

Deprint will prefix the output with `file:line` to make it easier to find the source of debug statements.

Output can be controlled by changing the global value `deprint.Output`. Possible values are:

- `Stdout` prints all statements os.Stdout (default)
- `Stderr` prints all statements os.Stderr
- `Disabled` disable all printing
