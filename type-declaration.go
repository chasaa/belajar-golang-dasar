package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var myKTP NoKTP = "111111"
	var marriedStatus Married = true

	fmt.Println(myKTP)
	fmt.Println(marriedStatus)
}
