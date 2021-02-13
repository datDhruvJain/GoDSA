package queues

import "testing"

func TestCircularQueue(t *testing.T) {
	c := New()
	*c.queue = make([]int, 5)
	c.enqueue(10)
	c.enqueue(20)
	c.enqueue(10)
	c.enqueue(20)
	c.enqueue(10)
	c.printinfo()
	c.enqueue(30)

	c.printinfo()
	c.printinfo()
	c.printinfo()
	c.printinfo()
	c.printinfo()
	c.printinfo()
	c.printinfo()
	c.printinfo()

	var p *[]int
	p = new([]int)
	*p = make([]int, 10)

}
