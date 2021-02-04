package Stacks

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	Push(1)
	Push(2)
	Push(3)
	Push(4)
	Push(5)
	if Top() != 5 {
		t.Log("The Top() function should have returned 5")
		t.Fail()
	}
	Push(6)
}

func TestPop(t *testing.T) {
	Pop()
	Pop()
	Push(6)
	Push(7)
	Push(8)
	lol := Pop()
	fmt.Print(lol)
	if lol != 7 {
		t.Log("Pop function should have returned 7")
		t.Fail()
	}
}
