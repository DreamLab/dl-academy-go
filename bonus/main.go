package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"
)

var pattern *regexp.Regexp // define a package-wide global variable, a pointer to type regexp.Regexp it can't be const, becuase it's evaluated at runtime and not compile time

func fetch(url string, wg *sync.WaitGroup) {
	// tell wait group that we have done. decrement goroutines counter
	defer wg.Done()

	// make HTTP request
	res, err := http.Get(url)
	if err != nil {
		// this is error that stop program
		fmt.Printf("Failed to fetch due to error %s %s\n", url, err.Error())
		return
	}

	if res.StatusCode != http.StatusOK { // we expect 200 OK
		fmt.Printf("Failed to fetch due to status code %s %d\n", url, res.StatusCode)
		return
	}

	fmt.Printf("Response for %s is %d \n", url, res.StatusCode)
	defer res.Body.Close()          // even if something goes wrong the body reader will still close
	fname := "tmp/" + safePath(url) // build a filename for the file
	f, err := os.Create(fname)      // create the file; note that os.Create does not create any directories along the way

	if err != nil { // check if the file has been succesfully created
		fmt.Println("Error creating file " + fname + " " + err.Error())
		return
	}

	defer f.Close()               // always defer as soon as possible to avoid trouble
	_, err = io.Copy(f, res.Body) // use io.Copy to move bytes from response body reader to file writer

	if err != nil { // io.Copy returns the number of bytes written and error, if any
		fmt.Println("Error copying response of" + url + " " + err.Error())
		return
	}

	fmt.Println("Fetched " + url + " as " + fname)

}

func safePath(url string) string { // define the safePath func to take a string and return a string
	return pattern.ReplaceAllString(url, "_") // replace all unsafe characters with underscore
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing required paramter - url list")
		fmt.Println("Usage:")
		fmt.Println("./main url.txt")
		os.Exit(2)
	}

	pattern = regexp.MustCompile("[<,>,:,\",/,\\,|,?,*]") // initialize our global variable, Regexp.MustCompile panics (Regexp.Compile returns an error) but that's better for safe initialization of our global variable
	// read data from file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("There was an error opening file", err)
		os.Exit(3)
	}

	// close file after end of current code block
	defer file.Close()

	// wait group for goroutines synchronization
	var wg sync.WaitGroup

	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// increment goroutines counter
		wg.Add(1)

		// run HTTP request in goroutines
		go fetch(scanner.Text(), &wg)
	}

	// wait for all goroutines to end
	wg.Wait()

	fmt.Println("End")
}
