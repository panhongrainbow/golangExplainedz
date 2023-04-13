package fake_interface_comparsion

import (
	"testing"
)

type DataFace interface {
	GetData() int
}

type MyData struct {
	data int
}

func (d *MyData) GetData() int {
	return d.data
}

func GetDataFunc() DataFace {
	return &MyData{}
}

var data DataFace

/*
or

func GetDataFunc() DataFace {
	return &MyData{}
}

var data DataFace

or

func GetDataFunc() MyData {
	return MyData{}
}

var data MyData

or

func GetDataFunc() *MyData {
	return &MyData{}
}

var data *MyData
*/

func Benchmark_interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data = GetDataFunc()
	}
}

/*
go tool pprof -http=:8080 mem.out
go test -bench=. -memprofile mem.out
*/
