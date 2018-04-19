package main

import "fmt"

func main () {
    // Make a map
    m := make(map[string]int)

    // Set
    m["k1"] = 7
    fmt.Println(m["k1"])

    // Get
    v := m["k1"]
    fmt.Println(v)

    // Delete
    delete(m, "k1")

    // Distinguish between empty values and non-existent ones
    if value, exists := m["k1"]; exists {
        fmt.Println("Value exists")
        fmt.Println(value)
    }
}