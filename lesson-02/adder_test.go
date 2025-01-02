package integers

import (
	"fmt"
	"testing"
)

// Example function used on documentations. You can run its tests by executing go test -v
// The comment with the output is validated, if you change to something else it will break the example
func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d received %d", expected, sum)
	}
}
