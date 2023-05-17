.PHONY: help

race: iocopy_race iocopy_fixed iocopy_enhanced
iocopy_race:
	go test -race -v -run Test_Race_iocopy ./race/ # | tail -n 3
iocopy_fixed:
	go test -race -v -run Test_Race_iocopy ./fixed/ | tail -n 3
iocopy_enhanced:
	go test -race -v -run Test_Race_iocopy ./enhanced/ | tail -n 3

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
	@echo "  iocopy_race            Run race condition test with iocopy"
	@echo "  iocopy_fixed           Run race condition test with fixed iocopy"
	@echo "  iocopy_enhanced        Run race condition test with atomic operations"
	@echo "  benchmark_fixed      Run benchmark for fixed race condition test"
	@echo "  benchmark_enhanced   Run benchmark for race condition test with atomic operations"
