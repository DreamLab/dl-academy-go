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
        total := 0 // keep track of the total job count
        jobs := make(chan string, 100) // make a channel for jobs
        results := make(chan string, 100) // and another one for their results

        for i := 1; i <= 3; i++ { // spawn 3 workers
            go fetchWorker(i, jobs, results) // every worker runs in a separate goroutine
        }

        for scanner.Scan() { // while there is a new line
            jobs <- scanner.Text() // send the line to the jobs channel, would block execution if channel becomes full
            total++ // increment the total number of jobs
        }
        close(jobs) // notify workers there's no more work by closing the channel, this does not erase any messages inside

        if total > 0 { // if there were any jobs
            for j := 1; j <= total; j++ { // wait for them to finish
                <-results // this blocks the executing until receiving one message, on each loop iteration - we wait for all jobs to finish
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

func fetchWorker (id int, jobs <-chan string, results chan<-string) { // define a worker that accepts its id, a read-only channel for jobs and a write-only channel for results
    fmt.Println("Worker", id, "started")
    for j := range jobs { // while the channel is open or has any messages
        fmt.Println("Worker", id, "starting job for", j)
        result := fetch(j) // fetch the resurce and save the result to variable
        fmt.Println("Worker", id, "finished job for", j)
        fmt.Println(result) // print the job result
        results <- result // send the result back to the main function, would block the execution if channel becomes full
    }
    fmt.Println("Worker", id, "finished") // this will be done once the for/while is finished, so when the channel is closed
}

func fetch (url string) (result string) { // the fetch function accepts an url as a string and returns a named string
    response, err := http.Get(url) // fetch the resource

    if err != nil { // if error occured on request
        result = "Failed to fetch due to error " + url + " " + err.Error()
    } else if response.StatusCode != http.StatusOK { // we expect 200 OK
        result = "Failed to fetch due to status code " + url + " " + fmt.Sprintf("%d", response.StatusCode) // response.StatusCode is an int and we don't rely on fmt.Println here so we need to turn it into a string
    } else { // process the response
        defer response.Body.Close() // even if something goes wrong the body reader will still close
        fname := "tmp/" + safePath(url) // build a filename for the file
        f, err := os.Create(fname) // create the file; note that os.Create does not create any directories along the way

        if err != nil { // check if the file has been succesfully created
            result = "Error creating file " + fname + " " + err.Error()
        } else {
            defer f.Close() // always defer as soon as possible to avoid trouble
            _, err := io.Copy(f, response.Body) // use io.Copy to move bytes from response body reader to file writer

            if err != nil { // io.Copy returns the number of bytes written and error, if any
                result = "Error copying response of" + url + " " + err.Error()
            } else { // otherwise we're done with this url
                result = "Fetched " + url + " as " + fname
            }
        }
    }
    return // every function returning something needs to explicitly return, this is done without a value becuase we are using a named return
} 