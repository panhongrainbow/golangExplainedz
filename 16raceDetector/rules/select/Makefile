.PHONY: help

test: select_race select_fixed
select_race:
	go test -race -v -run Test_Race_select ./race/ | tail -n 3
select_fixed:
	go test -race -v -run Test_Race_select ./fixed/ | tail -n 3
select_atomic:
	go test -race -v -run Test_Race_select ./atomic/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test            Run all race condition tests for all subtests"
	@echo "  sync_race       Run race condition test with sync mechanism"
	@echo "  sync_fixed      Run race condition test with fixed sync mechanism"
