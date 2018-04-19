package main

import "fmt"

func main () {
    // Make a slice
    s := make([]string, 3)

    // Set
    s[0] = "a"
    s[1] = "b"
    s[2] = "a"
    fmt.Println(s)

    // Get
    a := s[0]
    fmt.Println(a)

    // Length
    length := len(s)
    fmt.Println(length)

    // Append
    s = append(s, "b")
    fmt.Println(s)

    // Copy
    s2 := make([]string, len(s))
    copy(s2, s)
    fmt.Println(s2)

    // Slice
    l := s[2:3] // also :3 - up to 3 without 3; 2: - from 2 including 2
    fmt.Println(l)
    fmt.Println(s[:3])
    fmt.Println(s[2:])
}