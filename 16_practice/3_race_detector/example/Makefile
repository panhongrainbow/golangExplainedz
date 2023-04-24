.PHONY: help





# As long as the program contains `goroutine`
goroutine_race:
	go test -race -v -run Test_Check_data_race/global_variable_race ./
goroutine_no_race:
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
# As long as the program `accesses pointer variables`
pointer_race:
	go test -race -v -run Test_Check_data_race/pointer_race ./
pointer_no_race:
	go test -race -v -run Test_Check_data_race/pointer_no_race ./
# As long as the program `accesses pointer variables`
interface_race:
	go test -race -v -run Test_Check_data_race/interface_race ./
interface_no_race:
	go test -race -v -run Test_Check_data_race/interface_no_race ./
map_race:
	go test -race -v -run Test_Check_data_race/map_race ./
map_no_race:
	go test -race -v -run Test_Check_data_race/map_no_race ./
map_sync:
	go test -race -v -run Test_Check_data_race/map_sync ./
config_race:
	go test -race -v -run Test_Check_data_race/config_race ./
config_no_race:
	go test -race -v -run Test_Check_data_race/config_no_race ./
config_sync:
	go test -race -v -run Test_Check_data_race/config_sync ./
config_atomic:
	go test -race -v -run Test_Check_data_race/config_atomic ./
select_race:
	go test -race -v -run Test_Check_data_race/select_race ./
select_no_race:
	go test -race -v -run Test_Check_data_race/select_no_race ./
