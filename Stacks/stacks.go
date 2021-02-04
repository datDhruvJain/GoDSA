// TODo : add better documentation for stacks

/*
Remember : The variable t is a pointer AND a counter.
It counts the number of elements in the array and points the last element added to the array
*/
package Stacks

import (
	"fmt"
	"github.com/datdhruvjain/GoDSA/SearchingAlgorithms"
)

var stack = make([]int, 5)
var t int = -1

func Push(e int) {
	if t == len(stack)-1 {
		fmt.Println("Stack overflow error")
		return
	}

	t++
	stack[t] = e
}

func Pop() int {
	// here we use < -1 instead of != -1 so that if there is any error, we can recover out of it
	if t < -1 {
		// TODo :  add error handling, so that we dont have to return 0
		fmt.Println("Stack underflow error")
		return 0
	}

	temp := stack[t]
	t--
	return temp

	// alternatively, at the cost of being explicit,
	// return stack[t--]
}

func Top() int {
	if t > 0 {
		return stack[t]
	}
	return 0
}

func Search(x int) int {
	s := make([]int, len(stack)-1)
	copy(s, stack)
	return SearchingAlgorithms.BinarySearch(s, x)
}
