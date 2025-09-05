package main

import "fmt"

func main() {
	var slice []string
	slice = append(slice, "hello")
	slice = append(slice, "world")
	fmt.Printf("%T", cap(slice))
}
