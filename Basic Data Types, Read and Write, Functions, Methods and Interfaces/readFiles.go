package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	slice := make([]person, 0)
	var file string
	var text []string

	fmt.Print("Please enter the name of your text file(enter text.txt): ")
	fmt.Scan(&file)

	data, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	data.Close()

	for _, each_ln := range text {
		prs := person{fname: strings.Fields(each_ln)[0], lname: strings.Fields(each_ln)[1]}
		slice = append(slice, prs)
	}

	for _, element := range slice {
		fmt.Printf("fname: %s - lname: %s\n", element.fname, element.lname)
	}
}
