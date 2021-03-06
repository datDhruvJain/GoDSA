package sortingalgorithms

import (
	"fmt"
	"time"
)

// TODo : Write good documentation for this

// Remember : Arrays are passed by reference, not by value.
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			// len - i - 1, because the arr[len-1]th element has no arr[j+1], which we are using
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Printf("%v\r", arr)
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
	fmt.Println(arr)
}

/*
func main() {
	var arr []int = []int{5, 3, 1, 2, 6, 7, 9, 4}
	BubbleSort(arr)
}
*/
