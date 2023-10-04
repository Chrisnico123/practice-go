package main

import "fmt"

type Student struct{
	Name string
	Class string
}

func (s *Student) SetMyName(name string) {
	fmt.Println("Nama digantikan menjadi " , name)
	s.Name = name
}

func (s *Student) CallMyName() {
	fmt.Println("Hello, My Name is " , s.Name)
}

func main() {
	example := Student{Name: "Chrisnico",Class: "10"}

	example.SetMyName("Alexander")
	example.CallMyName()
}

