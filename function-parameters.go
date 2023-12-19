package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	// firstname:="Rizaldi"
	sayHelloTo("Rizaldi", "Yusuf")
	sayHelloTo("Asmikud", "Gaming")
}
