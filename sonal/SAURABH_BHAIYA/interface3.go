package main

import "fmt"

type student interface {
	studentName()
	studentRoll() int
	studentClass()
}

func (s studentDetails) studentName() {
	fmt.Println("student name is ", s.name)
}
func (s studentDetails) studentRoll() int {
	r := s.roll
	return r
}
func (s studentDetails) studentClass() {
	fmt.Println("student class", s.class)

}

type studentDetails struct {
	name  string
	roll  int
	class string
}

func main() {
	var st studentDetails
	st.name = "sonal"
	st.roll = 83
	st.class = "bca"
	var s student
	s = st

	s.studentName()
	fmt.Println("student roll number is ", s.studentRoll())
	s.studentClass()
}
