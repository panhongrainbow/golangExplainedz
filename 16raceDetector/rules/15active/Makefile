.PHONY: help

race: atomic_race atomic_fixed
atomic_race:
	go test -race -v -run Test_Race_atomic ./race/ # | tail -n 3
atomic_fixed:
	go test -race -v -run Test_Race_atomic ./fixed/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test                 Run all race condition tests for all subtests"
	@echo "  atomic_race            Run race condition test with atomic"
	@echo "  atomic_fixed           Run race condition test with fixed atomic"
