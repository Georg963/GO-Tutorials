package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	per   person
	class string
}

func (p *person) changeAge(newAge int) {
	p.age = newAge
}

func main() {
	per_1 := person{"George", 25}
	fmt.Println(per_1)

	//Methods
	fmt.Println(per_1.name, per_1.age)
	per_1.changeAge(26)
	fmt.Println(per_1)

	//Nested structs
	stu_1 := student{person{"George", 25}, "CS"}
	fmt.Println(stu_1)
	fmt.Println(stu_1.per.name, stu_1.per.age, stu_1.class)

	//change Age of stu_1
	stu_1.per.changeAge(26)
	fmt.Println(stu_1)

}
