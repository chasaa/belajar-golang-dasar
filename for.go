package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Chasa", "Bandung", "Indonesia"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Chasa", "Bandung", "Indonesia"}
	for _, name := range names {
		// fmt.Println("index", i, "=", name)
		fmt.Println("value:", name)
	}

	person := make(map[string]string)
	person["name"] = "Chasa"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println("key", key, "=", value)
	}
}
