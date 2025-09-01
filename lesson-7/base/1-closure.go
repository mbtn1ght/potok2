package main

import "fmt"

type Closure struct {
	i   *int
	str *string
}

func (c Closure) Func() {
	*(c.i)++
	*(c.str) = "hello"
}

func main() {
	var i int
	var str string

	closure := Closure{
		i:   &i,
		str: &str,
	}

	closure.Func()
	closure.Func()

	fmt.Println(i, str)
}
