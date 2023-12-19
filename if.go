package main

import (
	"fmt"
)

func main() {
	name := "Rizaldi"

	if name == "Rizal" {
		fmt.Println("Hello Rizal")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// var length = len(name)
	// if length > 5 {
	// 	fmt.Println("Terlalu panjang")
	// } else {
	// 	fmt.Println("Nama sudah cukup")
	// }

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah cukup")
	}
}
