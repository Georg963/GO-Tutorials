package main

import "fmt"

func main() {
	var student = map[string]int{"George": 26, "Björn": 24}
	//Adding Items
	student["Luise"] = 23
	fmt.Println(student)
	//Delete Items
	delete(student, "Luise")
	fmt.Println(student)
	fmt.Println()

	var person = make(map[string]int)
	person["Luise"] = 23
	person["Stephen"] = 24
	fmt.Println(person)

	//Iterate over a Map
	for key, element := range person {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
	fmt.Println()

	//merge Maps
	for k, v := range person {
		student[k] = v
	}

	fmt.Println(student)
	fmt.Println()

}
