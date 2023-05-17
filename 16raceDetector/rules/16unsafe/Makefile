.PHONY: help

race: unsafe_race unsafe_fixed unsafe_enhanced
unsafe_race:
	go test -race -v -run Test_Race_unsafe ./race/ | tail -n 3
unsafe_fixed:
	go test -race -v -run Test_Race_unsafe ./fixed/ | tail -n 3
unsafe_enhanced:
	go test -race -v -run Test_Race_unsafe ./enhanced/ | tail -n 3

benchmark: benchmark_fixed benchmark_enhanced
benchmark_fixed:
	go test -v -bench=. -run=none -benchmem ./fixed/
benchmark_enhanced:
	go test -v -bench=. -run=none -benchmem ./enhanced/

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  unsafe_race            Run race condition test with unsafe pointer"
	@echo "  unsafe_fixed           Run race condition test with fixed unsafe pointer"
	@echo "  unsafe_enhanced        Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_enhanced   Run benchmark for race condition test with atomic operations"