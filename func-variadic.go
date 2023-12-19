package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(100, 90, 80, 70, 60, 50)
	fmt.Println(total)

	slice := []int{100, 90, 80, 70, 60, 50}
	total = sumAll(slice...)
	fmt.Println(total)
}
