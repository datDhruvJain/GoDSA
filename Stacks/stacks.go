// TODo : add better documentation for stacks

/*
Package Stacks Implements Array based stacks and it's functions.
Remember : The variable t is a pointer AND a counter.
It counts the number of elements in the array and points the last element added to the array
*/
package Stacks

import (
	"fmt"
)

type IntStack struct {
	stack []int
	t     int
}

/*
New creates new stack of integer type and returns it's address (Also use pointers, Am I cool now? :) )
*/
func (is IntStack) New(size int) *IntStack {
	return &IntStack{make([]int, size), -1}
}

// Push Pushes an element onto the stack
func (is *IntStack) Push(e int) {
	if is.t == len(is.stack)-1 {
		fmt.Println("Stack overflow error")
		return
	}

	is.t++
	is.stack[is.t] = e
}

func (is *IntStack) Pop() int {
	// here we use < -1 instead of != -1 so that if there is any error, we can recover out of it
	if is.t < -1 {
		// TODo :  add error handling, so that we dont have to return 0
		fmt.Println("Stack underflow error")
		return 0
	}

	temp := is.stack[is.t]
	is.t--
	return temp

	// alternatively, at the cost of being explicit,
	// return stack[t--]
}

func (is *IntStack) Top() int {
	if is.t > 0 {
		return is.stack[is.t]
	}
	return 0
}
