.PHONY: help

race: select_race select_fixed select_enhanced
select_race:
	go test -race -v -run Test_Race_select ./race/ | tail -n 3
select_fixed:
	go test -race -v -run Test_Race_select ./fixed/ | tail -n 3
select_enhanced:
	go test -race -v -run Test_Race_select ./enhanced/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test              Run all race condition tests for all subtests"
	@echo "  select_race       Run race condition test with select mechanism"
	@echo "  select_fixed      Run race condition test with fixed select mechanism"
	@echo "  select_enhanced   Run race condition test with atomic operation"
