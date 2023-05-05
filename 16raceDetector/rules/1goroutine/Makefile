.PHONY: help

test: goroutine_race goroutine_fixed goroutine_atomic
goroutine_race:
	go test -race -v -run Test_Race_goroutines ./race/ | tail -n 3
goroutine_fixed:
	go test -race -v -run Test_Race_goroutines ./fixed/ | tail -n 3
goroutine_atomic:
	go test -race -v -run Test_Race_goroutines ./atomic/ | tail -n 3

benchmark: benchmark_fixed benchmark_atomic
benchmark_fixed:
	go test -v -bench=. -run=none -benchmem ./fixed/
benchmark_atomic:
	go test -v -bench=. -run=none -benchmem ./atomic/

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  goroutine_race       Run race condition test with goroutines"
	@echo "  goroutine_fixed      Run race condition test with fixed goroutines"
	@echo "  goroutine_atomic     Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_atomic     Run benchmark for race condition test with atomic operations"
