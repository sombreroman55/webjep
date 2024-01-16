BINARY_NAME=webjep

build:
	go build -o $(BINARY_NAME) *.go

run: build
	./$(BINARY_NAME)

test:
	go test -v ./...

clean:
	go clean
