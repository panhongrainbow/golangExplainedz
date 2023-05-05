.PHONY: help

test: timer_race timer_fixed timer_atomic
timer_race:
	go test -race -v -run Test_Race_timer ./race/ | tail -n 3
timer_fixed:
	go test -race -v -run Test_Race_timer ./fixed/ | tail -n 3
timer_atomic:
	go test -race -v -run Test_Race_timer ./atomic/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  timer_race       Run race condition test with timer"
	@echo "  timer_fixed      Run race condition test with fixed timer"
