OUTPUT_DIR=out
DOCS_DIR=docs
BINARY_NAME=main

clean:
	${info ****** Cleaning ******}
	go clean
	rm -rf $(OUTPUT_DIR)

gendocs:
	${info ****** Generating docs ******}
	rm -rf $(DOCS_DIR)
	swag init

install:
	${info ****** Installing dependencies ******}
	go mod tidy

build: install
	${info ****** Building ******}
	go build -o ./$(OUTPUT_DIR)/$(BINARY_NAME) -v

build-and-run: build
	${info ****** Building and running ******}
	./$(OUTPUT_DIR)/$(BINARY_NAME)

dev:
	${info ****** Running dev ******}
	go run main.go
