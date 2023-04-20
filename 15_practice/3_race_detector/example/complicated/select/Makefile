.PHONY: help

race:
	go test -race -v -run='^\QTest_Check' ./select_race/ | tail -n 3
fixed:
	go test -race -v -run='^\QTest_Check' ./select_fixed/ | tail -n 3
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