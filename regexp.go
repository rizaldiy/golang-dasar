package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("r([a-z]+)i")

	fmt.Println(regex.MatchString("rizaldi"))
	fmt.Println(regex.MatchString("rifaldi"))
	fmt.Println(regex.MatchString("ronAldi"))

	search := regex.FindAllString("rizaldi risaldi rifaldi ronaldi kiki asmikud", -1)
	fmt.Println(search)
}
