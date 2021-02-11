package main

import "fmt"

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

func main() {
	arr := new([]int)
	fmt.Println("Enter the size of array")
	arrSize := 0
	fmt.Scanln(&arrSize)

	*arr = make([]int, arrSize)
	*stack = make([]int, len(*arr))

	fmt.Println("Enter array elements")
	for i := 0; i < len(*arr); i++ {
		fmt.Scanln(&(*arr)[i])
	}

	for _, e := range *arr {
		push(e)
	}

	for i := 0; i < len(*arr); i++ {
		fmt.Println(pop())
	}
}
