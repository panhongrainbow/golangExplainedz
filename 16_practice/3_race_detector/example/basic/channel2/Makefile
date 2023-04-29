.PHONY: help

test: channel2_race channel2_fixed channel2_atomic
channel2_race:
	go test -race -v -run Test_Race_channel2 ./race/ | tail -n 3
channel2_fixed:
	go test -race -v -run Test_Race_channel2 ./fixed/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test                 Run all race condition tests for all subtests"
	@echo "  benchmark            Run all benchmarks for all subtests"
	@echo "  channel2_race       Run race condition test with channel"
	@echo "  channel2_fixed      Run race condition test with fixed channel"
