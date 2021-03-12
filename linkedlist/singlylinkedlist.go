package linkedlist

import (
	"errors"
	"fmt"
)

var count int = 0

type sllnode struct {
	data int
	next *sllnode
}

func Insert(head *sllnode, data int, index int) error {
	if index > count {
		return errors.New("attempting to insert at index " +
			"which does not exist for the given linked list")
	}

	if count == 0 && index == 0 {
		head = &sllnode{data, nil}
	}
	if head == nil {
		return errors.New("head of the linked list is nil")
	}

	var temp *sllnode = head

	for (*temp).next != nil {
		temp = (*temp).next
	}

	fmt.Printf("Adding new data element with data value: %v\n", data)
	(*temp).next = &sllnode{data, nil}

	return nil
}
