package main

import "fmt"

// takes an int value
func zeroval(ival int) {
    ival = 0
}

// takes a pointer to int value
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // &i returns the memory address of i
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // Pointers can be printed too.
    fmt.Println("pointer:", &i)
}
