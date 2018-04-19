package main

import "fmt"

// takes a read-only and a write-only channel
func pass (in <-chan string, out chan<- string) {
    txt := <- in
    out <- txt
}

func main() {
    // create bufferd channel
    in := make(chan string, 1)
    out := make(chan string, 1)

    // channel is buffered, no need for concurrent reader
    in <- "hello"
    pass(in, out)

    // message passed from channel to channel
    result := <- out
    fmt.Println(result)

    // channels can be closed
    close(in)
}
