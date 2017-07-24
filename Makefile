##This file builds golang binaries to the bin directory and also allows for cleaning them up.
all: darwin linux

darwin:
	env GOOS=darwin GOARCH=amd64 go build -o bin/fibweb .

linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/fibweb .

clean:
	rm -rf ./bin