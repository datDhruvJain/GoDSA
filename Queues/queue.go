/*
Queues follow a FILO: First in Last out Structure.

HEAD -> The counter/pointer that points to the starting of the data elements.
TAIL -> The counter/pointer that points to the end of the data elements. New elements are added to the tail

It has two operations, Enqueue, to add an element to the tail, and Dequeue.
Enqueue adds an element to the TAIL of the queue.
Dequeue removes and returns the HEAD element of the queue, and points HEAD to the next element.

Remember:
	1. The Initially, both head and tail point to the first element, ie q[0].
	2. The tail will most of the times point to an empty element.
	3. New elements are added to the tail.
*/
package Queues

import "fmt"

/*
type Queue struct {
	size int
	q  []int
	head int
	tail int
}

func NewQueue() *Queue {
	return &Queue{}
}



func (q *Queue) Q() []int {
	return q.q
}

func (q *Queue) SetQ(q []int) {
	q.q = q
}
*/

var size int = 5

var q = make([]int, size)
var head int = -1
var tail int = -1

/*
eg q[- - - - -]
The head will have value -1, and the tail will have value -1.

First enqueue operation:
q[1 - - - -]
Head still has value -1, tail has value q[0]

Adding more elements:
q[1 2 3 4 -]
Head still has value -1, tail has value q[3]
*/
func Enqueue(e int) {
	if tail != size-1 {
		tail++
		q[tail] = e
	} else {
		fmt.Printf("Array full, cannot add element %v\n", e)
	}
}

/*
eg q[- - - - -]
The head will have value -1, and the tail will have value -1.

First enqueue operation:
q[- 2 3 4 -]
Head has value 1 and points to q[1] = 2, tail has value q[3]

q[- - 3 4 -]
Head has value 2 and points to q[2] = 3, tail has value q[3]
*/
func Dequeue() int {
	if tail > 0 {
		head++
		temp := q[head]
		return temp
	} else {
		fmt.Println("Array emppty, cannot get any element, returning -1")
		return -1
	}
}

// Function to check the next element to take out
func Peek() int {
	if tail > 0 {
		return q[head+1]
	}
	return -1
}
