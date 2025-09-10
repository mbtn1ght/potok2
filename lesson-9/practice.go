// 1. Интерфейсы
// Задача:
// Создай интерфейс Speaker с методом Speak() string. Сделай две структуры:
// Dog и Cat, каждая должна реализовать Speak()
// (например, "Woof" и "Meow"). Напиши функцию SaySomething(s Speaker),
// которая выводит результат Speak().

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (s *Dog) Speak() string {
	return "Woof"
}

func (s *Cat) Speak() string {
	return "Meaw"
}

func SaySomething(s Speaker) {
	fmt.Println(s.Speak())
}
func main() {
	Cat := &Cat{}
	Dog := &Dog{}
	SaySomething(Cat)
	SaySomething(Dog)
}
