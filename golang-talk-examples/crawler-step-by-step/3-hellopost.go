package main // define package

import "fmt" // import necessary package to print to console
import "net/http"  // import part of the net package to serve http requests
import "io/ioutil" // import part of the io package to read from request readers

const addr = "localhost:55555" // define an address to listen on, this will not change so should be const

func main() { // define main function
    http.HandleFunc("/", handle) // add a handler to the default ServeMux
    err := http.ListenAndServe(addr, nil) // start listening on the addres and instruct to use the default ServeMux
    fmt.Println(err.Error()) // ListenAndServe blocks execution unless an error occurs, so we log that here
}

func handle (w http.ResponseWriter, r *http.Request) { // define a function that will handle requests
    fmt.Println("request from", r.RemoteAddr, "method", r.Method) // log an incoming request

    if r.Method == http.MethodPost { // we work with a body supplied by post request
        body, err := ioutil.ReadAll(r.Body)

        if err != nil {
            http.Error(w, "Error reading request body", http.StatusInternalServerError) // inform client about error
            return // http.Error doesn't end the request by itself
        }

        w.Header().Set("Content-Type", "text/plain")  // expose response header and set content-type
        w.WriteHeader(http.StatusOK)  // write the header to the outgoing socket with 200 status code
        fmt.Fprintf(w, "You said: ") // Start the response body
        fmt.Fprintf(w, string(body)) // convert body (a buffer) into a string and write it as the response body
        fmt.Fprintf(w, "\n") // end the response body with a newline character
    } else { // other HTTP methods
        w.WriteHeader(http.StatusMethodNotAllowed) // we don't accept anything other than POST
    }
}