package deprint

func ExamplePrintln() {
	Output = Stdout
	Println("Hello world")
	// Output: print_examples_test.go:5: Hello world
}

func ExamplePrintln_disabled() {
	Output = Disabled
	Println("Hello world")
	// Output:
}

func ExamplePrintf() {
	Output = Stdout
	Printf("Hello %s", "world")
	// Output: print_examples_test.go:17: Hello world
}

func ExamplePrintf_disabled() {
	Output = Disabled
	Printf("Hello %s", "world")
	// Output:
}
