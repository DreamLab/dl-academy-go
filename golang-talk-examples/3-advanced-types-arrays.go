package main

import "fmt"
    
func main () {
    // Declare an array
    var a[5] int

    // Set
    a[4] = 100
    fmt.Println(a)

    // Get
    b := a[4]
    fmt.Println(b)

    // Array length
    length := len(a)
    fmt.Println(length)

    // Declare and initialize
    aa := [5]int{1, 2, 3, 4, 5}
    fmt.Println(aa)

    // Declare a 2D array
    var twoD [2][3]int

    twoD[1][1] = 5
    fmt.Println(twoD)
}