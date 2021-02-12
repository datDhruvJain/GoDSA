package main

import (
	"fmt"
)

var errFlag bool = false
var stack *[]int = new([]int)
var t int = -1

func getSize() int {
	return t
}

func isEmpty() bool {
	if t == -1 {
		return true
	}
	return false
}

func isFull() bool {
	if t == getSize()-1 {
		return true
	}
	return false
}

func top() int {
	if !isEmpty() {
		return (*stack)[t]
	}
	return -1
}

func push(e int) {
	if !isFull() {
		t++
		(*stack)[t] = e
	}
}

func pop() int {
	if !isEmpty() {
		temp := (*stack)[t]
		t--
		return temp
	} else {
		errFlag = true
		return -1
	}
}

func postfix(arr []byte) {
	var a, b int
	for i := 0; i < len(arr); i++ {
		if !(arr[i] == '+' || arr[i] == '-' || arr[i] == '*' || arr[i] == '/' || arr[i] == '^') {
			// to convert from byte to int, substract byte '0' from the number byte
			push(int(arr[i] - '0'))
		} else if arr[i] == '+' {
			b = pop()
			a = pop()
			push(a + b)
		} else if arr[i] == '-' {
			b = pop()
			a = pop()
			push(a - b)
		} else if arr[i] == '*' {
			b = pop()
			a = pop()
			push(a * b)
		} else if arr[i] == '/' {
			b = pop()
			a = pop()
			push(a / b)
		} else if arr[i] == '^' {
			b = pop()
			a = pop()
			push(a ^ b)
		}
	}
}

func main() {
	var arrsize int

	fmt.Println("Enter size of array")
	fmt.Scanf("%v", &arrsize)
	arr := make([]byte, arrsize)

	fmt.Println("Enter array elements")
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%v", &arr[i])
	}
	fmt.Println(arr)
	// assign []int to "Value AT" the address stack pointer variable is holding
	*stack = make([]int, len(arr))

	postfix(arr)
	fmt.Println(pop())
}
