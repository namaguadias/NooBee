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

	student := Student{
		Name:  "Dhias Ulhaq Widjarnako",
		Class: "Intermediate",
	}

	student.CallMyName()

	student.SetMyName("Rincember")

	student.CallMyName()
}
