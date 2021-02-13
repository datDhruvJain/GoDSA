package queues

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
	got := Dequeue()
	want := 1
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
	got = Dequeue()
	want = 2
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
