# deprint

[![Go Reference](https://pkg.go.dev/badge/go.ptx.dk/deprint.svg)](https://pkg.go.dev/go.ptx.dk/deprint)
[![Github Workflow](https://github.com/ptxmac/deprint/actions/workflows/go.yml/badge.svg)](https://github.com/ptxmac/deprint/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/ptxmac/deprint/branch/master/graph/badge.svg)](https://codecov.io/gh/ptxmac/deprint)

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

### Controlling deprint from the environment

Calling `deprint.FromEnv` will automatically configure the output based on a selectable environment variable

```go
package sample

import "go.ptx.dk/deprint"

func init() {
	deprint.FromEnv("DEBUG", deprint.Disabled)
}
```

This will use the value of `DEBUG` to select the output. If `DEBUG` is not set, the 2nd argument will be used instead.
`DEBUG` can have the following values:

- 0 => Disable
- 1 => Use stdout
- 2 => USe stderr
