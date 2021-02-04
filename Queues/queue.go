/*
Queues follow a FILO: First in Last out Structure.

HEAD -> The counter/pointer that points to the starting of the data elements.
TAIL -> The counter/pointer that points to the end of the data elements.

eg q[- - - - -]
The head will point to q[0], and the tail will point to q[0] here

Adding one element:
q[1 - - - -]
The head will point to q[0] = 1, and the tail will point to q[1] = -

Adding more elements
q[1 2 3 4 -]
The head points to q[0] = 1, and the tail will point to q[4] = -

It has two operations, Enqueue, to add an element to the tail, and Dequeue.
Enqueue adds an element to the TAIL of the queue.
Dequeue removes and returns the HEAD element of the queue, and points HEAD to the next element.

Remember:
	1. The Initially, both head and tail point to the first element, ie q[0].
	2. The tail will most of the times point to an empty element
*/
package Queues

var q = make([]int, 5)
var head int = 0
var tail int = 0
var cnt int = 0

func Enqueue(e int) {
	if cnt != len(q) {
		q[tail] = e
		tail++
		cnt++
	}
}

func Dequeue() int {
	if cnt > 0 {
		temp := q[head]
		head++
		cnt--
		return temp
	}
	return 0
}
