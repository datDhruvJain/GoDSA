package Queues

import "testing"

func TestEnqueue(t *testing.T) {
	Enqueue(1)
	Enqueue(2)
	Enqueue(3)
	if Peek() != 1 {
		t.Log("Should be equal to 1")
		t.Fail()
	}
	Dequeue()
	if Peek() != 2 {
		t.Log("Should be equal to 2")
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	Enqueue(1)
	Enqueue(2)
	Enqueue(3)
	if Dequeue() != 1 {
		t.Log("Should be equal to 1")
		t.Fail()
	}
	if Dequeue() != 2 {
		t.Log("Should be equal to 2")
		t.Fail()
	}
}
