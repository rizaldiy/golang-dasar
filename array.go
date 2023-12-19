package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "rizal"
	names[1] = "yusuf"
	names[2] = "aldi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))

	var values1 = [...]int{
		99,
		88,
		77,
	}
	fmt.Println(values1)
	fmt.Println(len(values1))
	values1[0] = 100
	fmt.Println(values1)
}
