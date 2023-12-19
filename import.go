package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Rizal")
	// helper.sayGoodBye("Rizal") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
