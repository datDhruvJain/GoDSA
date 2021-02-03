package SearchingAlgorithms

import (
	"github.com/datdhruvjain/GoDSA/SortingAlgorithms"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var arr = []int{2, 5, 1, 6, 3, 8, 9, 0, 4}
	var x int = 9

	// Having a sorted array is necessary for Binary search to work
	SortingAlgorithms.BubbleSort(arr)

	if BinarySearch(arr, x) != 8 {
		t.Log("Array index returned should have been 8 for this test")
		t.Fail()
	}
}
