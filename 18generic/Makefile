.PHONY: help

test:
	go test -v -run='^Test_Check_' no_pointer_test.go
	@echo
	go test -v -run='^Test_Check_' pass_pointer_test.go
benchmark:
	go test -bench=. -run='^Benchmark_Performance_' no_pointer_test.go pass_pointer_test.go
group:
	go test -bench=. -run='^Benchmark_Group_' generic_test.go

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "This Makefile provides the following targets for testing and benchmarking:"
	@echo ""
	@echo "  test       - Run unit tests"
	@echo "               - Test_Check in no_pointer_test.go"
	@echo "               - Test_Check in pass_pointer_test.go"
	@echo ""
	@echo "  benchmark  - Run benchmark tests"
	@echo "               - Benchmark_Performance in no_pointer_test.go"
	@echo "                 and pass_pointer_test.go"
	@echo "               - Benchmark_Group in generic_test.go"
	@echo ""
	@echo "These targets are designed to ensure that the packages function correctly"
	@echo "and to identify any potential issues."
	@echo ""