package main // define package

import (
	"regexp" // we also need regexp library to sanitize filenames
	"net/http" // import part of the net package to serve http requests and make requests
	urlPkg "net/url"
	"fmt" // import necessary package to print to console
	"bufio" // we'll switch from io's reader to bufio's scanner to read the body line by line"os"
	"io" // but we'll still rely on io to copy from readers to writers
    "os" // we'll use it to create file writers
	"flag" // for parsing command line arguments
	"net" // for making tcp connection
	"time" // for seed for random string generator
	"compress/gzip" // gzip compress and uncompress
	"encoding/json" // encoding json
	"math/rand" // rand seed
	"strings" // parsing string
	"sync" // mutex for sync
)

var pattern *regexp.Regexp // define a package-wide global variable, a pointer to type regexp.Regexp it can't be const, becuase it's evaluated at runtime and not compile time


// function is executed on package require or it load
func init() {
	rand.Seed(time.Now().UnixNano())
}


// letters for random string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// generate random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// struct for storing crawl result
type resultData struct {
	url string
	data string
}


// single crawler job
type crawlerJob struct {
	id string
	urls []string
}

func NewCrawlerJob(urls []string) crawlerJob {
	return crawlerJob{RandStringRunes(8), urls}
}

var tasks chan crawlerJob // channel for crawler jobs
var finishedTasks map[string]map[string]string // map in which we store result of finished job
var inProgressTasks map[string]struct{} // map for easy lookup for enquued task
var httpsHosts map[string]struct{} // map of hosts that support HTTPS
var mapLock sync.RWMutex // lock for synchronization write and read from map between goroutines

// define HTTP client with configured timeouts
// Why? See https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
var httpClient = &http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 3 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}


func main() { // define main function
	var listenAddr string // in this var we store listen addr
	tasks = make(chan crawlerJob, 100)
	inProgressTasks = make(map[string]struct{})
	finishedTasks = make(map[string]map[string]string)
	httpsHosts = make(map[string]struct{})

	flag.StringVar(&listenAddr, "listen", "localhost:55555", "listen address for HTTP server" ) // user now can pass listen address to program
	flag.Parse()

	fmt.Println("Starting server on address", listenAddr)

	pattern = regexp.MustCompile("[<,>,:,\",/,\\,|,?,*]") // initialize our global variable, Regexp.MustCompile panics (Regexp.Compile returns an error) but that's better for safe initialization of our global variable
	http.HandleFunc("/", handle) // add a handler to the default ServeMux
	go tasksProcessor()
	err := http.ListenAndServe(listenAddr, nil) // start listening on the addres and instruct to use the default ServeMux
	fmt.Println(err.Error()) // ListenAndServe blocks execution unless an error occurs, so we log that here

}

func handle (w http.ResponseWriter, r *http.Request) { // define a function that will handle requests
	fmt.Println("request from", r.RemoteAddr, "method", r.Method) // log an incoming request

	if r.Method == http.MethodPost { // we work with a body supplied by post request
		scanner := bufio.NewScanner(r.Body) // create a new scanner instance
		urls := make([]string, 0)

		for scanner.Scan() { // while there is a new line
			urls = append(urls, scanner.Text())
		}

		task := NewCrawlerJob(urls)
		if err := scanner.Err(); err != nil { // scanner returns an error in a separate call
			http.Error(w, "Error reading request body", http.StatusInternalServerError) // inform client about error
			return // http.Error doesn't end the request by itself
		}

		tasks <- task
		mapLock.Lock()
		inProgressTasks[task.id] = struct{}{}
		mapLock.Unlock()

		w.WriteHeader(http.StatusCreated)  // write the header to the outgoing socket with 200 status code
		fmt.Fprintf(w, "job id: %s", task.id)
	} else if r.Method == http.MethodGet {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 1 {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return // http.Error doesn't end the request by itself
		}

		id := parts[1]
		fmt.Println("Get for job", id)

		mapLock.RLock()
		defer mapLock.RUnlock()
		if result, ok := finishedTasks[id]; ok {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)  // write the header to the outgoing socket with 200 status code
			reultStr, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return // http.Error doesn't end the request by itself
			}
			fmt.Fprint(w, string(reultStr))
			return
		} else if _, ok := inProgressTasks[id]; ok {
			w.WriteHeader(http.StatusOK)  // write the header to the outgoing socket with 200 status code
			fmt.Fprint(w, "Job in progress")
			return
		}

		http.Error(w, "job not found " + id, http.StatusNotFound)
		return // http.Error doesn't end the request by itself

	} else { // other HTTP methods
		w.WriteHeader(http.StatusMethodNotAllowed) // we don't accept anything other than POST
	}
	// response is automatically finished when the handle function returns
}

func safePath (url string) string { // define the safePath func to take a string and return a string
	return pattern.ReplaceAllString(url, "_") // replace all unsafe charaters with underscore
}

func fetchWorker (id int, jobs <-chan string, results chan<-resultData) { // define a worker that accepts its id, a read-only channel for jobs and a write-only channel for results
	fmt.Println("Worker", id, "started")
	for j := range jobs { // while the channel is open or has any messages
		fmt.Println("Worker", id, "starting job for", j)
		resultMsg := fetch(j) // fetch the resurce and save the result to variable
		fmt.Println("Worker", id, "finished job for", j)
		fmt.Println(resultMsg) // print the job result
		result := resultData{j, resultMsg}
		results <- result // send the result back to the main function, would block the execution if channel becomes full
	}
	fmt.Println("Worker", id, "finished") // this will be done once the for/while is finished, so when the channel is closed
}

func fetch (url string) (result string) { // the fetch function accepts an url as a string and returns a named string
	parsedUrl, err := urlPkg.Parse(url)
	if err != nil {
		return "Unable to parse url " + url + " err " + err.Error()
	}

	if parsedUrl.Scheme != "https" {
		if _, ok := httpsHosts[parsedUrl.Host]; ok {
			fmt.Printf("Changing url to HTTPS for %s (from map)\n", url)
            parsedUrl.Scheme = "https"
            url = parsedUrl.String()
		} else {
			conn, err := net.Dial("tcp", parsedUrl.Host + ":443")

			if err == nil {
				defer conn.Close()
				httpsHosts[parsedUrl.Host] = struct {}{}
				fmt.Printf("Changing url to HTTPS for %s\n", url)
                parsedUrl.Scheme = "https"
                url = parsedUrl.String()
			}
		}
	}


	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "Failed to create request due yo error " + url + " " + err.Error()
	}

	req.Header.Set("Accept-Encoding", "gzip")
	response, err := httpClient.Do(req) // fetch the resource

	if err != nil { // if error occured on request
		result = "Failed to fetch due to error " + url + " " + err.Error()
	} else if response.StatusCode != http.StatusOK { // we expect 200 OK
		result = "Failed to fetch due to status code " + url + " " + fmt.Sprintf("%d", response.StatusCode) // response.StatusCode is an int and we don't rely on fmt.Println here so we need to turn it into a string
	} else { // process the response
		defer response.Body.Close() // even if something goes wrong the body reader will still close
		// Check that the server actually sent compressed data
		var reader io.ReadCloser
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(response.Body)
			defer reader.Close()
		default:
			reader = response.Body
		}

		fname := "tmp/" + safePath(url) // build a filename for the file
		f, err := os.Create(fname) // create the file; note that os.Create does not create any directories along the way

		if err != nil { // check if the file has been succesfully created
			result = "Error creating file " + fname + " " + err.Error()
		} else {
			defer f.Close() // always defer as soon as possible to avoid trouble
			_, err := io.Copy(f, reader) // use io.Copy to move bytes from response body reader to file writer

			if err != nil { // io.Copy returns the number of bytes written and error, if any
				result = "Error copying response of" + url + " " + err.Error()
			} else { // otherwise we're done with this url
				result = "Fetched " + url + " as " + fname
			}
		}
	}
	return // every function returning something needs to explicitly return, this is done without a value becuase we are using a named return
}

func tasksProcessor() {
	for {
		select {
		case task := <-tasks:
			fmt.Println("Start processing task " + task.id)
			taskCount := len(task.urls)
			jobs := make(chan string, 100)        // make a channel for jobs
			results := make(chan resultData, 100) // and another one for their results
			report := make(map[string]string)

			for i := 1; i <= 3; i++ { // spawn 3 workers
				go fetchWorker(i, jobs, results) // every worker runs in a separate goroutine
			}

			for _, url := range task.urls {
				jobs <- url // send the line to the jobs channel, would block execution if channel becomes full
			}

			close(jobs) // notify workers there's no more work by closing the channel, this does not erase any messages inside

			for j := 0; j < taskCount; j++ { // wait for them to finish
				data := <-results // this blocks the executing until receiving one message, on each loop iteration - we wait for all jobs to finish
				report[data.url] = data.data
			}

			fmt.Println("End processing of task " + task.id)
			mapLock.Lock()
			finishedTasks[task.id] = report
			delete(inProgressTasks, task.id)
			mapLock.Unlock()
		default:

		}

	}
}
