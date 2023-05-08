.PHONY: help

race: interface_race interface_fixed interface_enhanced
interface_race:
	go test -race -v -run Test_Race_interface ./race/ | tail -n 3
interface_fixed:
	go test -race -v -run Test_fixed_interface ./race/ | tail -n 3
interface_enhanced:
	go test -race -v -run Test_atomic_interface ./race/

benchmark:
	go test -v -bench=. -run=none -benchmem ./race/

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  race                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  interface_race       Run race condition test with interface"
	@echo "  interface_fixed      Run race condition test with fixed interface"
	@echo "  interface_enhanced   Run race condition test with atomic operations"
