package main

import "fmt"
import "math"

func main () {
    // declaration
    var myVar string

    // assignment
    myVar = "Jak leci?"

    // declaration & assignment
    myOtherVar := "A spoko"
    fmt.Println(myVar, "\n", myOtherVar)

    // declaration & assignment of slice
    myTab := []string{"jak", "spoko", "to", "spoko"}
    fmt.Println(myTab)

    // constants
    const n = 100000000
    const d = 2e10 / n

    fmt.Println(int64(d))
    fmt.Println(math.Sin(d))
}