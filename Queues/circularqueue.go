/*
A Circular queue is a queue, in which the head and tail go round and round, not wasting any space
we need to look this up and write code for it
*/
package queues

import "fmt"

/*  front -> adds elements
	rear -> remove element from

	both are pointing to the present/current element

Full:
if front + 1 == rear, it is full,
also full if front = len -1 and rear ==0

Empty:
empty if both pointers are pointing to the same element */

type CircularQueue struct {
	queue *[]int

	// c.front points to the current element
	front int
	// rear points to the current element
	rear int
}

func New() *CircularQueue {
	//cq := CircularQueue{new([]int), -1, -1}
	var cq CircularQueue
	cq.queue = new([]int)
	cq.front = -1
	cq.rear = -1
	return &cq
}

func (c *CircularQueue) isFull() bool {
	if c.front+1 == c.rear || (c.front == len(*c.queue)-1 && c.rear == 0) {
		return true
	}
	return false
}

func (c *CircularQueue) isEmpty() bool {
	if c.front == c.rear {
		return true
	}
	return false
}

func (c *CircularQueue) enqueue(e int) {
	if !c.isFull() {
		c.front = (c.front + 1) % len(*c.queue)
		(*c.queue)[c.front] = e
		if c.rear == -1 {
			c.rear = 0
		}
	} else {
		fmt.Println("Stack overflow error")
	}
}

func (c *CircularQueue) dequeue() int {
	if !c.isEmpty() {
		c.rear = (c.rear + 1) % len(*c.queue)
		return (*c.queue)[c.rear]
	}
	return -1
}

// Helper function
func (c *CircularQueue) printinfo() {
	c.dequeue()
	fmt.Printf("The stack is %v. The c.c.front is %v and the rear is %v\n", *c.queue, c.front, c.rear)
}
