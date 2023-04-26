.PHONY: help

test: interface_race interface_fixed interface_atomic
interface_race:
	go test -race -v -run Test_Race_interface ./race/ | tail -n 3
interface_fixed:
	go test -race -v -run Test_fixed_interface ./race/ | tail -n 3
interface_atomic:
	go test -race -v -run Test_atomic_interface ./race/

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
	@echo "  interface_race       Run race condition test with interface"
	@echo "  interface_fixed      Run race condition test with fixed interface"
	@echo "  interface_atomic     Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_atomic     Run benchmark for race condition test with atomic operations"
