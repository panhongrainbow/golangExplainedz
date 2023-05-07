.PHONY: help

race: timer_race timer_fixed timer_enhanced
timer_race:
	go test -race -v -run Test_Race_timer ./race/ | tail -n 3
timer_fixed:
	go test -race -v -run Test_Race_timer ./fixed/ | tail -n 3
timer_enhanced:
	go test -race -v -run Test_Race_timer ./enhanced/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  race             Run all race condition tests for all subtests"
	@echo "  timer_race       Run race condition test with timer"
	@echo "  timer_fixed      Run race condition test with fixed timer"
	@echo "  timer_enhanced   Run race condition test with atomic operations"