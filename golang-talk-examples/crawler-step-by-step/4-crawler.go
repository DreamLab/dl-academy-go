package main // define package

import "fmt" // import necessary package to print to console
import "net/http"  // import part of the net package to serve http requests and make requests
import "bufio" // we'll switch from io's reader to bufio's scanner to read the body line by line
import "io" // but we'll still rely on io to copy from readers to writers
import "os" // we'll use it to create file writers
import "regexp" // we also need regexp library to sanitize filenames

const addr = "localhost:55555" // define an address to listen on, this will not change so should be const
var pattern *regexp.Regexp // define a package-wide global variable, a pointer to type regexp.Regexp it can't be const, becuase it's evaluated at runtime and not compile time

func main() { // define main function
    pattern = regexp.MustCompile("[<,>,:,\",/,\\,|,?,*]") // initialize our global variable, Regexp.MustCompile panics (Regexp.Compile returns an error) but that's better for safe initialization of our global variable 
    http.HandleFunc("/", handle) // add a handler to the default ServeMux
    err := http.ListenAndServe(addr, nil) // start listening on the addres and instruct to use the default ServeMux
    fmt.Println(err.Error()) // ListenAndServe blocks execution unless an error occurs, so we log that here
}

func handle (w http.ResponseWriter, r *http.Request) { // define a function that will handle requests
    fmt.Println("request from", r.RemoteAddr, "method", r.Method) // log an incoming request

    if r.Method == http.MethodPost { // we work with a body supplied by post request
        scanner := bufio.NewScanner(r.Body) // create a new scanner instance
        for scanner.Scan() { // while there is a new line
            url := scanner.Text() // save the line in a variable
            response, err := http.Get(url) // fetch the resource

            if err != nil { // if error occured on request
                fmt.Println("Failed to fetch due to error", url, err.Error()) // log it
            } else if response.StatusCode != http.StatusOK { // we expect 200 OK
                fmt.Println("Failed to fetch due to status code", url, response.StatusCode) // log it if it's not
            } else { // process the response
                defer response.Body.Close() // even if something goes wrong the body reader will still close
                fname := "tmp/" + safePath(url) // build a filename for the file
                f, err := os.Create(fname) // create the file; note that os.Create does not create any directories along the way

                if err != nil { // check if the file has been succesfully created
                    fmt.Println("Error creating file", fname, err.Error()) // print error if not
                } else {
                    defer f.Close() // always defer as soon as possible to avoid trouble
                    _, err := io.Copy(f, response.Body) // use io.Copy to move bytes from response body reader to file writer

                    if err != nil { // io.Copy returns the number of bytes written and error, if any
                        fmt.Println("Error copying response of", url, err.Error()) // print the error
                    } else { // otherwise we're done with this url
                        fmt.Println("Fetched", url, "as", fname) // print success message
                    }
                }
            }
        }

        if err := scanner.Err(); err != nil { // scanner returns an error in a separate call
            http.Error(w, "Error reading request body", http.StatusInternalServerError) // inform client about error
            return // http.Error doesn't end the request by itself
        }

        w.WriteHeader(http.StatusOK)  // write the header to the outgoing socket with 200 status code
    } else { // other HTTP methods
        w.WriteHeader(http.StatusMethodNotAllowed) // we don't accept anything other than POST
    }
    // response is automatically finished when the handle function returns
}

func safePath (url string) string { // define the safePath func to take a string and return a string
    return pattern.ReplaceAllString(url, "_") // replace all unsafe charaters with underscore 
}