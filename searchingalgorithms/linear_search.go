package searchingalgorithms

/*
Linear Search is a brute force method to search for an element in an array. It
will search throughout the array, element by element until it finds the element
to be searched. Time complexity of Linear search is O(n), that is, it is a
linear time algorithm.
*/
func LinearSearch(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}

	return -1
}
