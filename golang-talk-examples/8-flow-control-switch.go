package main

import "fmt"

func main () {
    // Simple version
    i := 2
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("something different")
    }

    // You can combine multiple cases
    switch i {
    case 1, 2:
        fmt.Println("one or two")
    default:
        fmt.Println("something different")
    }
}