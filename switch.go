package main

import "fmt"

func main() {
	name := "Chasa"

	switch name {
	case "Chasa":
		fmt.Println("Hello Chasa")
		fmt.Println("Hello Chasa")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hello, boleh kenalan?")
		fmt.Println("Hello, boleh kenalan?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
