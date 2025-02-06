package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d but got %d", expected, sum)
	}
}

// The output comment will also be compiled
// Try to remove and add it back while running the "go test -v" command
func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}