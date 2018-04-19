// prepare a report
//     write any errors and success information to the response writer for the client to see
// make the listening port configurable
//     have a look at the flags package
// gzip
//     optimize the crawling further by sending "Accept-Encoding: gzip" header with your requests
//     use the "compress/gzip" package to uncompress the body before saving
//     do the same the other way - compress the report in response if client accepts gzip
// asynchronous crawling
//     instead of keeping the client waiting return a StatusAccepted reply with an url the client can use to query for the job status
//     provide him with a progress about the number of URLs failed/completed/total
//     after crawling is finished return a report on that url
// try to implement a HTTPS everywhere functionality
//     parse the url to determine hostname and schema
//     if it's HTTP, try to do a TCP Connect to port 443 of the target hostname to see if it can be reached over HTTPS
//     if yes, try to fetch the file over HTTPS
//     if no, or the above fails, fallback to HTTP
//     remember the result in a map for future requests to the same hostname

package main

import "fmt"

func main() {
    fmt.Println("Have fun!")
}