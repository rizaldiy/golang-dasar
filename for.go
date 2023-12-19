package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Rizal", "Yusuf", "Aldi", "Kiki", "Devara", "Asmikud"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Rizal"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}
}
