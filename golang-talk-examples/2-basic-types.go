package main

import "fmt"

func main () {
    // Strings and chars
    stringVar := "String"
    charVar := 'a'
    fmt.Println(stringVar, charVar, string(charVar))

    // Integers
    var a int // also int8 int16 int32 int64
    var b uint // also uint8 uint16 uint32 uint64
    simpleInt := 5
    fmt.Println(a, b, simpleInt)

    // floats
    var c float32 // also float64
    simpleFloat := 5.5
    fmt.Println(c, simpleFloat)

    // special
    var d byte // alias for uint8
    var e rune // alias for int32
    var f complex128 // also complex64 part of math/cmplx
    fmt.Println(d, e, f)
}