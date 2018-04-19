# go tool pprof
Built-in go proflier

# Prepare data for profiler
```
go test   -bench=.  -benchmem -cpuprofile prof.cpu -memprofile prof.mem
```

# Basic command in profiler
* top10 - top 10 allocation or % of CPU function
* list FUN_NAME - source code of function with information about CPU/Memory consumption 

# CPU profiler
```
go tool pprof pprof.test  prof.cpu
```
# Memory profier
```
go tool pprof -alloc_objects pprof.test prof.mem
```
