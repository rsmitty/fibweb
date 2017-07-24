## FibWeb

A toy webserver in golang that returns Fibonacci sequences.

### Motivation

This toy was created to familiarize myself with building webservices in golang. Additionally, I wanted to learn more about creating tests for services that I build.

### Building

This client can be easily built using the provided Makefile. Simply issue one of the following:
- `make darwin` - Builds 64-bit Mac client in the bin/ directory
- `make linux` - Builds 64-bit Linux client in the bin/ directory
- `make all` - Builds both of the above in the bin/ directory

### Usage and Assumptions

This small binary assumes that you will run it from a terminal simply as a test. It will accept requests on port 8080 of localhost in your environment.

- To start the binary, simply issue `./bin/fibweb` from the root of the fibweb directory.
- Once the server is running, you can then interact with it via curl (in a different terminal window) or in the browser at `http://localhost:8080`

There are some "rules" around interacting with the server:
- The server expects to receive requests on the `/fib` endpoint. If a different endpoint is provided, an error will be returned directing you to `http://localhost:8080/fib`
- The `/fib` endpoint expects a variable called `COUNT` to be provided with the request. This variable contains the desired length of your Fibonacci sequence. Thus a request will look something like `curl http://localhost:8080?COUNT=5`.
- The `COUNT` value supplied must be a positive number. Anything <= 0 will result in an error message being returned.
- If the count value is too high, it will overflow a 64-bit integer. Thus, a submitted `COUNT` of greater than 93 will return an error.

### Testing
There are tests written for this toy server that use the built-in golang testing and httptest packages. Tests can be run by simply issuing `go test` from the root of this directory.

### Examples
Examples of interacting with the fibweb server binary.

Normal request:
```bash
$ curl localhost:8080/fib?COUNT=5

[0 1 1 2 3]
```

Incorrect path:
```bash
$ curl localhost:8080/incorrectpath?COUNT=5

ERROR: Path not recognized. Use /fib.
```

Incorrect input:
```bash
$ curl localhost:8080/fib?FOO=5

ERROR: Unable to retrieve COUNT input. Please try again.
```

Integer overflow:
```bash
$ curl localhost:8080/fib?COUNT=94

ERROR: This calculation would cause a 64bit integer overflow. Choose a value <= 93
```