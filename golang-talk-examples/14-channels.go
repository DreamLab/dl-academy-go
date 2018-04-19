package main

import "fmt"

func main () {
    // create channel
    messages := make(chan string)

    // write to channel, this is blocking operation so to work we should do it in other thread
    go func () {     
        messages <- "elo"
    }()

    // receive data from channel
    txt := <- messages
    fmt.Println(txt)
}