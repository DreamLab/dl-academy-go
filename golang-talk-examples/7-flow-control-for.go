package main

import "fmt"

func main () {
    // Simple version
    i := 1
    for i <= 3 {
        i = i + 1
    }

    // Classic for-loop
    for j := 1; j <= 9; j++ {
        if j == 2 {
            continue
        } else if j == 5 {
            break
        }
        fmt.Println(j)
    }

    // For without condition loops until break
    for {
        fmt.Println("loop")
        break
    }
}