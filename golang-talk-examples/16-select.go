package main

import "fmt"

func main() {
    // create unbuffered channels
    ch1 := make(chan string)
    ch2 := make(chan string)

    // no messages so default is executed
    select {
    case msg := <- ch1:
        fmt.Println("msg received on ch1", msg)
    case msg := <- ch2:
        fmt.Println("msg received on ch2", msg)
    default:
        fmt.Println("no messages")
    }
    // no listener attached so default is executed
    select {
    case ch1 <- "hello":
        fmt.Println("message sent")
    default:
        fmt.Println("failed to send")
    }
}
