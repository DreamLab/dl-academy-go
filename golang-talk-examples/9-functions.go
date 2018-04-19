package main

import "fmt"

// Single return function
func times(a int, b int) int {
    return a * b
}

// Named returns
func threeSum (a, b, c int) (d int) {
   d = a + b
   d = d + c
   return
}

// Multiple returns
func ThreePrimes() (int, int, int) {
   return 2, 3, 5
}

func main () {
    first := times(2, 3)
    second := threeSum(1, 2, 3)
    thirdOne, _, thirdThree := ThreePrimes()

    fmt.Println("2 times 3 is", first)
    fmt.Println("sum of 1, 2 and 3 is", second)
    fmt.Println("The first and the third prime numbers are", thirdOne, thirdThree)
}