package iteration

import (
	"fmt"
	"testing"
)

// Benchmark tests can be run with `go test --bench=.`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func BenchmarkRepeatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatBuilder("a", 4)
	}
}

func ExampleRepeat() {
	text := Repeat("a", 4)
	fmt.Print(text)
	// Output: aaaa
}

func ExampleRepeatBuilder() {
	text := RepeatBuilder("a", 4)
	fmt.Print(text)
	// Output: aaaa
}

func TestRepeating(t *testing.T) {
	given := Repeat("a", 4)
	expected := "aaaa"

	if given != expected {
		t.Errorf("Expected %q but received %q", expected, given)
	}
}
