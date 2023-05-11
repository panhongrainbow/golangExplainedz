.PHONY: help

hello:
	go build -o hello.out hello.go
	objdump -S hello.out > hello.txt

clean:
	rm -f *.out
	rm -f *.txt

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Compile, disassemble, clean up files."
	@echo ""
	@echo "  make hello - compile hello.go into hello.out and generate assembly in hello.txt"
	@echo "  make clean - delete .out and .txt files"
	@echo "  make help - print this help message"