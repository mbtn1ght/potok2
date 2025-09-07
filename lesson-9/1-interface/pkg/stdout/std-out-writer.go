package main

import (
	"fmt"
)

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (s *Writer) Write(p []byte) (int, error) {
	return fmt.Println(string(p))
}

func main() {
	w := NewWriter()
	slice := []byte("Hallo world!")
	w.Write(slice)
}
