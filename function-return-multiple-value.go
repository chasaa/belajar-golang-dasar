package main

import "fmt"

func getFullName() (string, string, string) {
	return "Chasa", "Werkudoro", "Indonesia"
}

func main() {
	firstName, middleName, _ := getFullName()
	fmt.Println(firstName, middleName)
}
