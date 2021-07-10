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
	resetOutput()
	out, outErr := captureOutput(t, func() {
		Println("Hello")
	})
	expect := "print_test.go:23: Hello\n"
	if out != expect {
		t.Errorf("Expected: [%s], got: [%s]", expect, out)
	}
	if outErr != "" {
		t.Errorf("Expected no stderr output, got: %s", outErr)
	}
}

func TestPrintln_Disabled(t *testing.T) {
	Output = Disabled
	defer resetOutput()
	out, outErr := captureOutput(t, func() {
		Println("Hello")
	})
	if out != "" {
		t.Errorf("Expected no output, got: %s", out)
	}
	if outErr != "" {
		t.Errorf("Expected no stderr output, got: %s", outErr)
	}
}

func TestPrintln_Stderr(t *testing.T) {
	Output = Stderr
	defer resetOutput()
	out, outErr := captureOutput(t, func() {
		Println("Hello")
	})
	expect := "print_test.go:52: Hello\n"
	if out != "" {
		t.Errorf("Expected no output, got: %s", out)
	}
	if outErr != expect {
		t.Errorf("Expected: [%s], got: [%s]", expect, outErr)
	}
}

func TestPrintln_fromEnv(t *testing.T) {
	tests := []struct {
		name  string
		value string
		out   bool
		err   bool
	}{
		{
			name:  "disabled",
			value: "0",
		},
		{
			name:  "stdout",
			value: "1",
			out:   true,
		},
		{
			name:  "stderr",
			value: "2",
			err:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer withEnv(t, "TEST_DEBUG", tt.value)()
			defer resetOutput()

			FromEnv("TEST_DEBUG", Stdout)

			out, outErr := captureOutput(t, func() {
				Println("test fromenv")
			})
			expect := "print_test.go:93: test fromenv\n"
			if tt.out {
				if out != expect {
					t.Errorf("Expected: [%s], got: [%s]", expect, out)
				}
			} else {
				if out != "" {
					t.Errorf("Expected no output, got: %s", out)
				}
			}
			if tt.err {
				if outErr != expect {
					t.Errorf("Expected: [%s], got: [%s]", expect, outErr)
				}
			} else {
				if outErr != "" {
					t.Errorf("Expected no output, got: %s", outErr)
				}
			}
		})
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

// captureOutput replaces os.Std{out,err} while running f and returns the outputs as a strings
func captureOutput(t *testing.T, f func()) (string, string) {
	t.Helper()
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	defer func() {
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	}()

	or, ow, err := os.Pipe()
	noError(t, err)
	er, ew, err := os.Pipe()
	noError(t, err)

	os.Stdout = ow
	os.Stderr = ew

	f()

	noError(t, ow.Close())
	noError(t, ew.Close())
	bsOut, err := ioutil.ReadAll(or)
	noError(t, err)
	bsErr, err := ioutil.ReadAll(er)
	noError(t, err)

	noError(t, or.Close())
	noError(t, er.Close())

	return string(bsOut), string(bsErr)
}

// withEnv sets key=value and returns a function to reset it back
func withEnv(t *testing.T, key string, value string) func() {
	t.Helper()
	old, ok := os.LookupEnv(key)
	noError(t, os.Setenv(key, value))
	return func() {
		if ok {
			noError(t, os.Setenv(key, old))
		} else {
			noError(t, os.Unsetenv(key))
		}
	}
}

// resetOutput sets output the to default value
func resetOutput() {
	Output = Stdout
}
