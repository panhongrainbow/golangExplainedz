.PHONY: help

race: closure_race closure_fixed closure_enhanced
closure_race:
	go test -race -v -run Test_Race_closure ./race/ | tail -n 3
closure_fixed:
	go test -race -v -run Test_Race_closure ./fixed/ | tail -n 3
closure_enhanced:
	go test -race -v -run Test_Race_closure ./enhanced/ | tail -n 3

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  race                 Run all race condition tests for all subtests"
	@echo "  closure_race         Run race condition test with closure"
	@echo "  closure_fixed        Run race condition test with fixed closure"
	@echo "  interface_enhanced   Run race condition test with atomic operations"
