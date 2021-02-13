package searchingalgorithms

/*
The time complexity for this Search is O(log2(n)).
This is because each time the search cannot find the element, the size of the array halves.

For example, if the size of the array is 32, the size reduces if the element is not found in the parent array, and it becomes 16.
this repeats (32, 16, 8, 4, 2, 1). Here then we require 6 guesses at most and is in the relation log2(n), as compared to n guesses.

The limitation here is that the array needs to be sorted before the operation can be applied

Points to Remember:
	1. There are three variables, low, high, mid
	  1.1 The low variable initially has value 0 and thus points to value at arr[0]
	  1.2 The high variable initially has value len(arr) - 1. This points to the last
	  element in the array
	  1.3 mid = (low + high)/2
	  1.4 this works because mid, high
	  and low values are ints and are floored in case of floating point answers
	2. The for loop will run until high >= low
	This in esscence means, until high does not pass low, the loop should keep running
	It does not work in any other way.
	3. returning
	  3.1 if the mid value is equal to the search value, return mid
	  3.2 if the mid value is lower than the search value, assign high to mid - 1
	  3.3 if the mid value is high that the search value, assign low to mid + 1
	  3.4 if the value is not found, return 0
*/
func BinarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	var mid int
	for high >= low {
		mid = (high + low) / 2

		if arr[mid] == x {
			return mid
		} else if x > arr[mid] {
			low = mid + 1
		} else if x < arr[mid] {
			high = mid - 1
		}
	}
	return -1
}
