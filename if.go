package main

import "fmt"

func main() {
	name := "Chasa Agung"

	if name == "Chasa" {
		fmt.Println("Hello Chasa")
	} else if name == "Agung" {
		fmt.Println("Hello Agung")
	} else {
		fmt.Println("Hi, Boleh kenalan ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
