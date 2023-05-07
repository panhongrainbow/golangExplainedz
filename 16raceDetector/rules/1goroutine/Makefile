.PHONY: help

race: goroutine_race goroutine_fixed goroutine_enhanced
goroutine_race:
	go test -race -v -run Test_Race_goroutines ./race/ | tail -n 3
goroutine_fixed:
	go test -race -v -run Test_Race_goroutines ./fixed/ | tail -n 3
goroutine_enhanced:
	go test -race -v -run Test_Race_goroutines ./enhanced/ | tail -n 3

benchmark: benchmark_fixed benchmark_enhanced
benchmark_fixed:
	go test -v -bench=. -run=none -benchmem ./fixed/
benchmark_enhanced:
	go test -v -bench=. -run=none -benchmem ./enhanced/

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  race                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  goroutine_race       Run race condition test with goroutines"
	@echo "  goroutine_fixed      Run race condition test with fixed goroutines"
	@echo "  goroutine_enhance    Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_enhanced   Run benchmark for race condition test with atomic operations"
