package linkedlist

import (
	"testing"
)

var head *sllnode

func TestInsert(t *testing.T) {
	e := Insert(head, 12, 0)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
}
