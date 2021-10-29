package main

import "fmt"

/*Interface describes all the methods of a method set and provides the signatures for each method.
https://www.golangprograms.com/go-language/interface.html
*/

type person struct {
	name string
	age  int
}

type student struct {
	name string
	age  int
}

func (p *person) changeAge(newAge int) {
	p.age = newAge
}

func (p *student) changeAge(newAge int) {
	p.age = newAge
}

type changeA interface {
	changeAge(newAge int)
}

func change(ca changeA, x int) {
	ca.changeAge(x)
}

func addOne(x interface{}) interface{} {
	switch x.(type) {
	case int:
		return x.(int) + 1
	case float64:
		return x.(float64) + 1
	default:
		return "Unknown Type"
	}
}

func main() {
	//Polymorphism and Type assertion
	per_1 := person{"George", 25}
	stu_1 := student{"George", 25}
	fmt.Println(per_1)
	fmt.Println(stu_1)

	per_1.changeAge(23)
	stu_1.changeAge(22)

	fmt.Println(per_1)
	fmt.Println(stu_1)

	change(&per_1, 26)
	change(&stu_1, 27)

	fmt.Println(per_1)
	fmt.Println(stu_1)

	//another way to use interfaces
	var c changeA = &stu_1
	change(c, 20)
	c.changeAge(21)

	fmt.Println(stu_1)

	var int_1 interface{}
	int_1 = 5
	var float_1 interface{}
	float_1 = 6.0

	fmt.Println(addOne(int_1))
	fmt.Println(addOne(float_1))
}
