package SearchingAlgorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	var arr = []int{2, 5, 1, 6, 3, 8, 9, 0, 4}
	var x int = 9

	BinarySearch(arr, x)
}
