.PHONY: help

test: channel_race channel_fixed
channel_race:
	go test -race -v -run Test_Race_channel ./race/ | tail -n 3
channel_fixed:
	go test -race -v -run Test_Race_channel ./fixed/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test               Run all race condition tests for all subtests"
	@echo "  channel_race       Run race condition test with channel"
	@echo "  channel_fixed      Run race condition test with fixed channel"
