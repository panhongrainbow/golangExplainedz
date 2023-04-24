.PHONY: help

test: map_race map_fixed map_atomic
map_race:
	go test -race -v -run Test_Race_map ./race/ | tail -n 3
map_fixed:
	go test -race -v -run Test_Race_map ./fixed/ | tail -n 3
map_atomic:
	go test -race -v -run Test_Race_map ./atomic/ | tail -n 3

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
	@echo "  map_race       Run race condition test with map"
	@echo "  map_fixed      Run race condition test with fixed map"
	@echo "  map_atomic     Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_atomic     Run benchmark for race condition test with atomic operations"
