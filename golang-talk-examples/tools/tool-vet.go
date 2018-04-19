package main

import "fmt"

func main() {
    str := "hello world!"
    // wrong format
    fmt.Printf("%d\n", str)

    // wrong type
    fmt.Printf("%s\n", &str)
}
