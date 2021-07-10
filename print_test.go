package deprint

import (
	"io/ioutil"
	"os"
	"testing"
)

func ExamplePrintln() {
	Println("Hello world")
	// Output: print_test.go:10: Hello world
}

func ExamplePrintln_disabled() {
	Output = Disabled
	Println("Hello world")
	// Output:
}

func TestPrintln(t *testing.T) {
	out := captureStdout(t, func() {
		Println("Hello")
	})
	expect := "print_test.go:22: Hello\n"
	if out != expect {
		t.Errorf("Expected: [%s], got: [%s]", expect, out)
	}
}

func TestPrintln_Disabled(t *testing.T) {
	Output = Disabled
	defer func() {
		Output = Stdout
	}()
	out := captureStdout(t, func() {
		Println("Hello")
	})
	if out != "" {
		t.Errorf("Expected no output, got: %s", out)
	}
}

func TestPrintln_Stderr(t *testing.T) {
	Output = Stderr
	defer func() {
		Output = Stdout
	}()
	out := captureStderr(t, func() {
		Println("Hello")
	})
	expect := "print_test.go:49: Hello\n"
	if out != expect {
		t.Errorf("Expected: [%s], got: [%s]", expect, out)
	}
}

// noError fails the test if err is not nil
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// Helpers

// captureStdout captures os.Stdout
func captureStdout(t *testing.T, f func()) string {
	return captureOutput(t, &os.Stdout, f)
}

// captureStderr captures os.Stderr
func captureStderr(t *testing.T, f func()) string {
	return captureOutput(t, &os.Stderr, f)
}

// captureOutput replaces target while running f and returns the output as a string
func captureOutput(t *testing.T, target **os.File, f func()) string {
	t.Helper()
	oldStdout := *target
	defer func() {
		*target = oldStdout
	}()

	r, w, err := os.Pipe()
	noError(t, err)

	*target = w
	f()
	noError(t, w.Close())
	bs, err := ioutil.ReadAll(r)
	noError(t, err)
	noError(t, r.Close())
	return string(bs)
}
