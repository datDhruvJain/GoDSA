package main

import "fmt"

var errFlag bool = false
var stack *[]byte = new([]byte)
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

func top() byte {
	if !isEmpty() {
		return (*stack)[t]
	}
	return 0
}

func push(e byte) {
	if !isFull() {
		t++
		(*stack)[t] = e
	}
}

func pop() byte {
	if !isEmpty() {
		temp := (*stack)[t]
		t--
		return temp
	} else {
		errFlag = true
		return 0
	}
}

func checkBalance(arr []byte) {
	var a byte
	for i := 0; i < len(arr); i++ {
		if arr[i] == '[' || arr[i] == '{' || arr[i] == '(' {
			push(arr[i])
		} else if arr[i] == ')' {
			a = pop()
			if a != '(' {
				fmt.Println("Unbalanced")
			}
		} else if arr[i] == '}' {
			a = pop()
			if a != '{' {
				fmt.Println("Unbalanced")
			}
		} else if arr[i] == ']' {
			a = pop()
			if a != '[' {
				fmt.Println("Unbalanced")
			}
		}

	}
	if !isEmpty() {
		fmt.Println("Unbalanced")
	}
}

func main() {
	arr := []byte{'(', '1', '2', '{', '}', ')'}
	//arr := []byte{'(', '1', '2', '{', '}',}
	*stack = make([]byte, len(arr))
	checkBalance(arr)
}
