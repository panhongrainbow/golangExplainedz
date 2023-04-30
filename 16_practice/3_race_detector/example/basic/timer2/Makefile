.PHONY: help

test: timer2_race timer2_fixed timer2_atomic
timer2_race:
	go test -race -v -run Test_Race_timer2 ./race/ | tail -n 3
timer2_fixed:
	go test -race -v -run Test_Race_timer2 ./fixed/ | tail -n 3
timer2_atomic:
	go test -race -v -run Test_Race_timer2 ./atomic/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  timer2_race       Run race condition test with timer"
	@echo "  timer2_fixed      Run race condition test with fixed timer"
