package main

import "fmt"

type CustomValue struct {
    value int
}

func (v CustomValue) GetValue () int {
    return v.value
}

func main () {
    a := 5

    // Simple version
    if a == 2 {
        fmt.Println("variable is equal 2")
    } else {
        fmt.Println("variable is NOT equal 2")
    }

    // Cascading
    if a == 2 {
        fmt.Println("variable is equal 2")
    } else if a == 3 {
        fmt.Println("variable is equal 3")
    } else {
        fmt.Println("variable is equal neither 2 nor 3")
    }

    b := CustomValue{5}

    // Declaration in statement
    if c := b.GetValue(); c < 5 {
        fmt.Println("b's value is less than 5")
    }
    // c no longer in scope
    // fmt.Println(c)
}