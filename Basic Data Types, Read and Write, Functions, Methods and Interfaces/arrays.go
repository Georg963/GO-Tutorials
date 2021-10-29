package main

import "fmt"

func main() {
	var myArr [5]int
	myArr[0] = 1
	myArr[1] = 2
	myArr[2] = 3
	myArr[3] = 4
	myArr[4] = 5

	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}
	fmt.Println()

	myArr_2 := [5]int{5, 4, 3, 2, 1}
	for i := 0; i < len(myArr_2); i++ {
		fmt.Println(myArr[i])
	}
	fmt.Println()

	//slices
	var mySlice []int
	for i := 0; i < 5; i++ {
		mySlice = append(mySlice, i)
	}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
	fmt.Println()

	//using the keyword make
	mySlice_2 := make([]int, 10) //when length and capacity is same
	for i := 0; i < len(mySlice_2); i++ {
		mySlice_2[i] = i
	}

	for i := 0; i < len(mySlice_2); i++ {
		fmt.Println(mySlice_2[i])
	}

	//Append a slice to an existing slice
	var cities_1 = []string{"Berlin", "München", "Potsdam"}
	var cities_2 = []string{"Dortmund", "Ulm"}
	fmt.Println(append(cities_1, cities_2...))

}
