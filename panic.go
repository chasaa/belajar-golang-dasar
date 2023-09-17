package main

import "fmt"

func endApp() {
    message := recover()
    if message != nil {
        fmt.Println("Error dengan message:", message)
    }
    fmt.Println("Aplikasi Selesai")
}

func runApp(err bool) {
    defer endApp()
    if err {
        panic("APLIKASI ERROR")
    }
    fmt.Println("Aplikasi Berjalan")
}

func main() {
    runApp(true)   
    fmt.Println("INITIALIZATION")
}