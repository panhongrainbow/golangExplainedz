.PHONY: help

test: sync_race sync_fixed
sync_race:
	go test -race -v -run Test_Race_sync ./race/ | tail -n 3
sync_fixed:
	go test -race -v -run Test_Race_sync ./fixed/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test            Run all race condition tests for all subtests"
	@echo "  sync_race       Run race condition test with sync mechanism"
	@echo "  sync_fixed      Run race condition test with fixed sync mechanism"
