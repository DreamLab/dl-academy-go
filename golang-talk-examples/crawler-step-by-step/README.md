# Simple example of crawler
HTTP server that download given urls and save it to file

## Setup 
```bash
mkdir src # in programs cwd
```
## 1-helloworld
Hello world in Go

##  2-hellohttp
Hello world in HTTP server. Prepare HTTP server that respond with 'Hello world'

Run HTTP server

```bash
go run 2-hellohttp.go
```

Test it

```bash
curl -X GET localhost:55555 -i
```

## 3-hellopost
HTTP server that print body of POST

Run server

```bash
go run 3-hellopost.go
```

Test it

```bash
curl -X POST localhost:55555 -i -d 'Eloszka'
```

## 4-crawler and 5-crawlerworkers

Create crawler that download given urls. Crawler from 5-crawlerworkers uses 3 goroutines workers to execute tasks.

```bash
go run 5-crawlerworkers.go
```

Test it

```bash
curl -X POST localhost:55555 -i -d 'https://www.onet.pl/
https://www.dreamlab.pl/
http://ocdn.eu/pulscms-transforms/1/Wc0ktkqTURBXy9lMGM4NjA4NzUxZmJhNWZiYWFkYzI5OTY0NTFmOGVlNC5qcGVnkpUDWQDNFlTNDuOTBV87 
'
```

or from file

```bash
$ cat urls.txt
https://www.onet.pl/
https://www.dreamlab.pl/
http://ocdn.eu/pulscms-transforms/1/Wc0ktkqTURBXy9lMGM4NjA4NzUxZmJhNWZiYWFkYzI5OTY0NTFmOGVlNC5qcGVnkpUDWQDNFlTNDuOTBV87 

$ curl -X POST localhost:55555 -i --data-binary @urls.txt

```

## 6-whatsnext

What to do next
