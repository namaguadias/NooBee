package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s *Student) CallMyName() {
	fmt.Printf("Hello, My Name is %s.\n", s.Name)
}

func main() {

	s := Student{
		Name:  "Dhias Ulhaq Widjarnako",
		Class: "Intermediate",
	}

	s.CallMyName()

	s.SetMyName("Rincember")

	s.CallMyName()
}
