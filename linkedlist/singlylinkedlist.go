package linkedlist

import (
	"errors"
	"fmt"
)

type sllnode struct {
	data int
	node *sllnode
}

func Insert(head *sllnode) (int, error) {
	if head == nil {
		return -1, errors.New("Head of the linked list is nil")
	}
}
