.PHONY: help

race: list_race list_fixed list_enhanced
list_race:
	go test -race -v -run Test_Race_list ./race/ | tail -n 3
list_fixed:
	go test -race -v -run Test_Race_list ./fixed/ | tail -n 3
list_enhanced:
	go test -race -v -run Test_Race_list ./enhanced/ | tail -n 3

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
	@echo "  list_race            Run race condition test with list"
	@echo "  list_fixed           Run race condition test with fixed list"
	@echo "  list_enhanced        Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_enhanced   Run benchmark for race condition test with atomic operations"
