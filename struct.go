package main

import "fmt"

type Customer struct {
    Name, Address string
    Age int
}

func main() {
    var chasa Customer
    chasa.Name = "Chasa"
    chasa.Address = "Indonesia"
    chasa.Age = 30

    fmt.Println(chasa.Name)
    fmt.Println(chasa.Address)
    fmt.Println(chasa.Age)

    joko := Customer{
        Name: "Joko",
        Address: "Cirebon",
        Age: 35,
    }
    fmt.Println(joko)
}