package main

import "fmt"

func getFullName1() (firstName, middleName, lastName string) { //func getFullName1() (firstName string, middleName string, lastName string)
	firstName = "Rizal"
	middleName = " "
	lastName = "Yusuf"

	return
}

func main() {
	a, b, c := getFullName1()
	fmt.Println(a + b + c)
}

