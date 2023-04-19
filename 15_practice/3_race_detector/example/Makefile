.PHONY: help

global_race:
	go test -race -v -run Test_Check_data_race/global_variable_race ./
global_no_race:
	go test -race -v -run Test_Check_data_race/global_variable_no_race ./
channel_race:
	go test -race -v -run Test_Check_data_race/channel_race ./
channel_no_race:
	go test -race -v -run Test_Check_data_race/channel_no_race ./
sync_race:
	go test -race -v -run Test_Check_data_race/sync_race ./
sync_no_race:
	go test -race -v -run Test_Check_data_race/sync_no_race ./
closure_race:
	go test -race -v -run Test_Check_data_race/sync_race ./
closure_no_race:
	go test -race -v -run Test_Check_data_race/sync_no_race ./
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "This mechanism is a suite of tests designed to ensure that"
	@echo "the packages are functioning correctly and"
	@echo "to identify any issues that may exist."
	@echo ""
	@echo "Available targets:"
	@echo "  test     - unit test"
	@echo "  cover    - coverage test"
	@echo "  race     - race test"
	@echo ""