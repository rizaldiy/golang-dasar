package main

import "fmt"

func main() {
	name := "Rizal"
	counter := 0

	increment := func() {
		name := "Yusuf"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
