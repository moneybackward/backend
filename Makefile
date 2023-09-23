BINARY_NAME=main

clean:
	${info ****** Cleaning ******}
	go clean

install:
	${info ****** Installing dependencies ******}
	go mod tidy

build: install
	${info ****** Building ******}
	go build -o $(BINARY_NAME) -v

build-and-run: build
	${info ****** Building and running ******}
	./$(BINARY_NAME)

dev:
	${info ****** Running dev ******}
	go run main.go
