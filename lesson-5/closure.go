package main

import "fmt"

func main() {

	intSeq := func() func() int {
		i := 0
		return func() int { // Это замыкание
			i++
			return i
		}
	}

	nextInt := intSeq()    // Возвращается замыкание
	fmt.Println(nextInt()) // Выведет 1
	fmt.Println(nextInt()) // Выведет 2
}
