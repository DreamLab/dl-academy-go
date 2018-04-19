package main // define package

import "fmt" // import necessary package to print to console
import "net/http"  // import part of the net package to serve http requests

const addr = "localhost:55555" // define an address to listen on, this will not change so should be const

func main() { // define main function
    http.HandleFunc("/", handle) // add a handler to the default ServeMux
    err := http.ListenAndServe(addr, nil) // start listening on the addres and instruct to use the default ServeMux
    fmt.Println(err.Error()) // ListenAndServe blocks execution unless an error occurs, so we log that here
}

func handle (w http.ResponseWriter, r *http.Request) { // define a function that will handle requests
    fmt.Println("hello to", r.RemoteAddr) // log an incoming request
    w.Header().Set("Content-Type", "text/plain")  // expose response header and set content-type
    w.WriteHeader(http.StatusOK)  // write the header to the outgoing socket with 200 status code
    _, err := fmt.Fprintf(w, "hello world\n") // use Fprintf to wrap writing a string to the ResponseWriter

    if err != nil { // fmt.Fprintf (and fmt.Println!) returns the number of bytes written (we don't care) and an error, if any
        fmt.Println(err.Error()) // print the error messsage of the error object
    }
    // response is automatically finished when the handle function returns
}