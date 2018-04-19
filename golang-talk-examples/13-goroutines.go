package main

import "fmt"
import "time"

// this task takes 3 seconds to complete
func longTask(txt string) {
    for i := 0; i < 3; i++ {
        fmt.Println(txt, ":", i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    fmt.Println("hello from main thread")
    longTask("hello from longTask")
    fmt.Println("longTask has ended")

    // we start a new goroutine
    go longTask("hello from longTask")

    // this gets printed right away
    fmt.Println("longTask might still be on")

    // wait for user input
    fmt.Scanln()
}