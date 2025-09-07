package main

import "fmt"

type Counter struct {
	count int
}

func NewCounter(start int) *Counter {
	return &Counter{start}
}

func (s *Counter) Increment() {
	s.count++
}

func (s *Counter) Value() int {
	return s.count
}

func main() {
	counter := NewCounter(10)
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.Value())

}
