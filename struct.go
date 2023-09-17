package main

import "fmt"

type Customer struct {
    Name, Address string
    Age int
}

func (customer Customer) sayHi(name string) {
    fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (customer Customer) sayHuuu() {
    fmt.Println("Huuu from", customer.Name)
}

func main() {
    var chasa Customer
    chasa.Name = "Chasa"
    chasa.Address = "Indonesia"
    chasa.Age = 30

    chasa.sayHi("Joko")
    chasa.sayHuuu()

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