package main

import (
	"fmt"
	"sort"
)

//Bubble Sort
func sortList(l []int, c chan []int) {
	var tmp int
	sliceLength := len(l)
	for i := 0; i < sliceLength; i++ {
		for j := 0; j < sliceLength-i-1; j++ {
			if l[j] > l[j+1] {
				tmp = l[j+1]
				l[j+1] = l[j]
				l[j] = tmp
			}
		}
	}
	c <- l
}

func main() {

	var intSlice []int

	fmt.Println("Enter your ints (-101 to exit): ")
	for {
		var input int
		fmt.Scan(&input)

		if input == -101 {
			break
		} else {
			intSlice = append(intSlice, input)
		}

	}

	middle := len(intSlice) / 2
	q1 := middle / 2
	q3 := middle + q1

	ch_1 := make(chan []int, 4)
	go sortList(intSlice[:q1+1], ch_1)
	go sortList(intSlice[q1+1:middle+1], ch_1)
	go sortList(intSlice[middle+1:q3+1], ch_1)
	go sortList(intSlice[q3+1:], ch_1)
	part_1 := <-ch_1
	part_2 := <-ch_1
	part_3 := <-ch_1
	part_4 := <-ch_1

	close(ch_1)

	fmt.Println("Part_1: ", part_1)
	fmt.Println("Part_2: ", part_2)
	fmt.Println("Part_3: ", part_3)
	fmt.Println("Part_4: ", part_4)

	//merge the 4 sorted subarrays into one large sorted array.
	part_5 := append(part_1, part_2...)
	part_6 := append(part_3, part_4...)
	sortedList := append(part_5, part_6...)
	sort.Ints(sortedList)

	fmt.Println("Sorted List: ", sortedList)
}
