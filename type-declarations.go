package main

import "fmt"

func main() {
	type NoKTP string
	type isMarried bool

	var noKtpRzl NoKTP = "3175061106970019"
	var marriedStatus isMarried = false

	fmt.Println(noKtpRzl)
	fmt.Println(marriedStatus)
}
