package packageDemo

import "testing"

// Functions begin with Benchmark
// This benchmark will run when tests are all passed and run with
//     go test -bench .
func BenchmarkDemo(b *testing.B) {
	for i := 0; i < 10000000; i++ {
		x := 1 + 1
		if x > 2 {
			x = 3
		}
	}
}

// Functions begin with Test
func TestDemo(t *testing.T) {
	if Demo() == -1 {
		t.Log("This will not be printed")
		t.Fail() // not executed
	}
}

func TestDemoWithFailure(t *testing.T) {
	if DemoWithFailure() == -1 {
		t.Log("should return 0")
		t.Fail() // will not stop here, continue
	}
}

func TestDemoWithFailureTwice(t *testing.T) {
	if DemoWithFailure() == -1 {
		t.Log("should return 0")
		t.FailNow() // test will terminate
	}
}

func TestDemoWithFailureThird(t *testing.T) { // will not be executed cause t.FailNow is called before
	if DemoWithFailure() == -1 {
		t.Log("should return 0")
		t.FailNow() // test will terminate
	}
}

// Function begin with Example lead to two results
// 1. go test will varify if the last two line of comments will generate same result of previous lines
// 2. The example will be printed in godoc
func ExampleDemo() {
	Demo() // This will print "This is a demo"
	// Output:
	// This is a demo
}

func ExampleDemoWithFailure() {
	DemoWithFailure() // This will print "This is a failed demo"
	// Output:
	// This is a failed demo
}
