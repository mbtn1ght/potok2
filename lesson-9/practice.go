// Задача:
// Создай интерфейс Mover с методом Move() string.
// Сделай две структуры: Car и Bird. Каждая должна реализовать
// Move(), например, "Car drives" и "Bird flies". Напиши функцию
// DoMove(m Mover), которая выводит результат Move().
package main

import "fmt"

type Move interface {
	Move() string
}

type Car struct{}
type Bird struct{}

func (s *Car) Move() string {
	return "Car drives"
}

func (s *Bird) Move() string {
	return "Bird flies"
}

func DoMove(s Move) {
	fmt.Println(s.Move())
}

func main() {
	Car := &Car{}
	Bird := &Bird{}
	DoMove(Car)
	DoMove(Bird)
}
