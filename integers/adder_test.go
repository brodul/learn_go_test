package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	// This is actualy tested via testing framework
	sum := Add(5, 1)
	// Output: 6
	fmt.Println(sum)
}
