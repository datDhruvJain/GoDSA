package SearchingAlgorithms

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	var arr = []int{2, 5, 1, 6, 3, 8, 9, 0, 4}
	var x int = 9

	if LinearSearch(arr, x) != 6 {
		t.Log("The index returned Should be 6")
		t.Fail()
	}
}

func ExampleLinearSearch() {
	var arr = []int{2, 5, 1, 6, 3, 8, 9, 0, 4}
	var x int = 9

	fmt.Print(LinearSearch(arr, x))
	// Output: 6
}
