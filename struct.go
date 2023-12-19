package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	var rizal Customer
	rizal.Name = "Rizal"
	rizal.Address = "Jakarta"
	rizal.Age = 26

	rizal.sayHi("Asmikud")
	rizal.sayHuuu()
	// fmt.Println(rizal.Name)
	// fmt.Println(rizal.Address)
	// fmt.Println(rizal.Age)

	// kiki := Customer{
	// 	Name:    "kiki",
	// 	Address: "Jakarta",
	// 	Age:     25,
	// }
	// fmt.Println(kiki)

	// budi := Customer{"budi", "Bandung", 26}
	// fmt.Println(budi)
}
