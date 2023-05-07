.PHONY: help

race: sync2_race sync2_fixed
sync2_race:
	go test -race -v -run Test_Race_sync2 ./race/ | tail -n 3
sync2_fixed:
	go test -race -v -run Test_Race_sync2 ./fixed/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  race             Run all race condition tests for all subtests"
	@echo "  sync2_race       Run race condition test with sync.Mutex"
	@echo "  sync2_fixed      Run race condition test with fixed sync.Mutex"
