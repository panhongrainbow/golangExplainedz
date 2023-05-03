.PHONY: help

test:
	go test -v ./

build:
	go test -c -o test ./

clean:
	rm -f test

help:
	@echo "Usage: make [target]"
	@echo "This Makefile is used to manage test, build, run and clean process of the project."
	@echo "It contains the following targets:"
	@echo
	@echo "  test   - Run tests"
	@echo "  build  - Compile the program and generate executable file test"
	@echo "  clean  - Delete the compiled test file"
	@echo