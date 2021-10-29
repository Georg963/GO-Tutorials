package main

import (
	"OOP/students"
	"fmt"
)

type person struct {
	student students.Student
	address string
}

func main() {
	var student_1 students.Student
	student_1.Init("Elias", 22, "TUM")
	student_1.GetInfo()

	student_1.ChangeAge(23)
	student_1.ChangeSchool("TU Berlin")
	student_1.GetInfo()
	fmt.Println("-----")
	person_1 := person{student_1, "Berlin"}
	fmt.Println(person_1)
	person_1.student.GetInfo()
	fmt.Printf("Address: %v\n", person_1.address)
}
