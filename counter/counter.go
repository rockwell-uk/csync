package counter

type Operation int

const (
	ADD Operation = iota
)

type Counter struct {
	ops   chan Operation
	res   chan int
	count int
}

func New() *Counter {
	counter := Counter{
		ops: make(chan Operation, 1000000),
		res: make(chan int, 1),
	}

	go func() {
		for range counter.ops {
			counter.res <- counter.count
			counter.count++
		}
	}()

	return &counter
}

func (c *Counter) Add() int {
	c.ops <- ADD
	return <-c.res
}
