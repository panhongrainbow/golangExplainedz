.PHONY: help

run: run_enhanced
run_enhanced:
	go test -v -run Test_Race_map ./enhanced/

race: map_race map_fixed map_enhanced
map_race:
	go test -race -v -run Test_Race_map ./race/ | tail -n 3
map_fixed:
	go test -race -v -run Test_Race_map ./fixed/ | tail -n 3
map_enhanced:
	go test -race -v -run Test_Race_map ./enhanced/ | tail -n 3

benchmark: benchmark_fixed benchmark_enhanced
benchmark_fixed:
	go test -v -bench=. -run=none -benchmem ./fixed/
benchmark_enhanced:
	go test -v -bench=. -run=none -benchmem ./enhanced/

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run                  Run some tests"
	@echo "  race                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  map_race             Run race condition test with map"
	@echo "  map_fixed            Run race condition test with fixed map"
	@echo "  map_atomic           Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_enhanced   Run benchmark for race condition test with atomic operations"
