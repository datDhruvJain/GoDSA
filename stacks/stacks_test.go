package stacks

import (
	"fmt"
	"testing"
)

var ist IntStack

func TestNew(t *testing.T) {
	ist.New(10)
	fmt.Println(len(ist.stack))
}
