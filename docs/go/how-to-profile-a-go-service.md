# What is profiling?

> In software engineering, profiling ("program profiling", "software profiling") is a form of dynamic program analysis that measures, for example, the space (memory) or time complexity of a program, the usage of particular instructions, or the frequency and duration of function calls. Most commonly, profiling information serves to aid program optimization, and more specifically, performance engineering. Profiling is achieved by instrumenting either the program source code or its binary executable form using a tool called a profiler (or code profiler). Profilers may use a number of different techniques, such as event-based, statistical, instrumented, and simulation methods.

[Wiki Reference](https://en.wikipedia.org/wiki/Profiling_(computer_programming))

**TL;DR:** It's a process to analyze the time, space, and instruction complexity of a program in an effort to improve it's performace or to understand it's limits.

## Why do we care?

- Memory usage
- Instruction complexity
- Program optimization
- Go routine usage
- Cost of compute/memory when running the application
- Reliability of application

## Golang Tooling

Go provides a few utilities that aid in this process.

- [pprof](https://pkg.go.dev/runtime/pprof) - The runtime profiling tool that can be added to any golang program.
- [net pprof](https://pkg.go.dev/net/http/pprof) - An http interface that can be added to an http server creating api's for an end user to access pprof data of a remote server.
- [go tool](https://pkg.go.dev/cmd/go) - The go tool-kit that is used to interact with and digest pprof data.

## How do I access pprof output?

Each service that uses the `chassis` module will run `pprof_net` when in `dev` mode.

**Endpoints**:

```txt
GET    /debug/pprof/
GET    /debug/pprof/heap
GET    /debug/pprof/goroutine
GET    /debug/pprof/allocs
GET    /debug/pprof/block
GET    /debug/pprof/threadcreate
GET    /debug/pprof/cmdline
GET    /debug/pprof/profile
GET    /debug/pprof/symbol
POST   /debug/pprof/symbol
GET    /debug/pprof/trace
GET    /debug/pprof/mutex
```

## Example

The below steps assume you have access to the `galactus` repo, `make`, `go` and `docker` installed on your local machine.

### Setup your local environment

```sh
# run all core infra locally in docker (mongo, pg, rabbitmq, proxy),
# and the core services of the whole system (command, eventer, notifier).
make local

# run the notifier load test, this will establish 1000 connections to the notifier service simulating 1000 connected users.
make notifier-load-test TEST_COUNT=1000
```

### Review heap usage

```sh
# run go tool connecting to the notifier service
go tool pprof http://localhost:8086/debug/pprof/heap

# pprof will enter interactive mode. (type help to see all available commands.)
help

# top - will show total allocations on the heap.
top

# example output
Type: inuse_space
Time: Dec 9, 2021 at 9:57am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 10355.50kB, 100% of 10355.50kB total
Showing top 10 nodes out of 57
      flat  flat%   sum%        cum   cum%
 4100.50kB 39.60% 39.60%  4100.50kB 39.60%  runtime.allocm
 1577.92kB 15.24% 54.83%  1577.92kB 15.24%  github.com/denisenkom/go-mssqldb/internal/cp.init
 1024.41kB  9.89% 64.73%  1024.41kB  9.89%  runtime.malg
  544.67kB  5.26% 69.99%   544.67kB  5.26%  github.com/xdg/stringprep.init
  544.67kB  5.26% 75.25%   544.67kB  5.26%  google.golang.org/grpc/internal/transport.newBufWriter
     514kB  4.96% 80.21%      514kB  4.96%  hash/crc32.archInitCastagnoli
  513.31kB  4.96% 85.17%   513.31kB  4.96%  regexp/syntax.(*compiler).inst
  512.01kB  4.94% 90.11%   512.01kB  4.94%  regexp/syntax.appendRange
  512.01kB  4.94% 95.06%   512.01kB  4.94%  sync.(*Map).Store
     512kB  4.94%   100%      512kB  4.94%  runtime.doaddtimer

# png - will output graph of total usage
png

# view the file
open profile001.png
```

### Review go routines

```sh
# review the goroutines that are running in the service
go tool pprof http://localhost:8086/debug/pprof/goroutine

# view top output
top

# example output
Type: goroutine
Time: Dec 9, 2021 at 10:42am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1017, 99.80% of 1019 total
Dropped 56 nodes (cum <= 5)
Showing top 10 nodes out of 14
      flat  flat%   sum%        cum   cum%
      1017 99.80% 99.80%       1017 99.80%  runtime.gopark
         0     0% 99.80%       1000 98.14%  atlas%2enotifier._Notifier_Connect_Handler
         0     0% 99.80%       1000 98.14%  google.golang.org/grpc.(*Server).handleStream
         0     0% 99.80%       1000 98.14%  google.golang.org/grpc.(*Server).processStreamingRPC
         0     0% 99.80%       1000 98.14%  google.golang.org/grpc.(*Server).serveStreams.func1.2
         0     0% 99.80%       1000 98.14%  gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc.StreamServerInterceptor.func1
         0     0% 99.80%          6  0.59%  internal/poll.(*pollDesc).wait
         0     0% 99.80%          6  0.59%  internal/poll.(*pollDesc).waitRead
         0     0% 99.80%          6  0.59%  internal/poll.runtime_pollWait
         0     0% 99.80%       1000 98.14%  notifier/handler.(*notifierHandler).Connect

# create the graph
png

# view the graph
open profile002.png
```

### System cleanup

```sh
# cleanup all local processes
make clean-local
```

### References

- [Mux integration](https://newbedev.com/profiling-go-web-application-built-with-gorilla-s-mux-with-net-http-pprof)
- [Gist](https://gist.github.com/slok/33dad1d0d0bae07977e6d32bcc010188)
- [example](https://blog.intelligentbee.com/2017/08/01/profiling-web-applications-golang/)
