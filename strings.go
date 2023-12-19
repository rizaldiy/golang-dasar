package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rizaldi Yusuf", "Rizal"))
	fmt.Println(strings.Contains("Rizaldi Yusuf", "asmikud"))

	fmt.Println(strings.Split("Rizaldi Yusuf", " "))

	fmt.Println(strings.ToLower("Rizaldi YUSUF"))
	fmt.Println(strings.ToUpper("rizaldi yusuf"))
	fmt.Println(strings.ToTitle("rizaldi yusuf"))

	fmt.Println(strings.Trim("          Rizaldi Yusuf    ", " "))
	fmt.Println(strings.ReplaceAll("Asmikud Asmikud", "Asmikud", "Rizaldi"))
}
