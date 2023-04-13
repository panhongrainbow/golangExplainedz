package main

import (
	"fmt"
	"sync"
	"testing"
)

var _ DataFace = (*MyData)(nil)
var _ DataFace = (*MyData2)(nil)
var _ DataFace = (*MyData3)(nil)

type DataFace interface {
	GetData() int
}

type MyData struct {
	data int
	Lock sync.Mutex
}

func (d *MyData) GetData() int {
	fmt.Println("In MyData")
	return d.data
}

type MyData2 MyData

func (d *MyData2) GetData() int {
	fmt.Println("In MyData2")
	return d.data
}

type MyData3 MyData

func (d *MyData3) GetData() int {
	fmt.Println("In MyData3")
	return d.data
}

func TransferByInterface() DataFace {
	data := MyData2{}
	return &data
}

/*
func TransferByInterface() DataFace {
	data := MyData3{}
	return &data
}
*/

func TransferByFakeInterface(data MyData) MyData3 {
	return MyData3(data)
}

/*
or
func TransferByFakeInterface(data MyData) MyData3 {
	return MyData3(data)
}
*/

var data MyData2

func Test_fake_interface(t *testing.T) {
	byInterface := TransferByInterface()
	byInterface.GetData()
	var data MyData
	byFakeInterface := TransferByFakeInterface(data)
	byFakeInterface.GetData()
}
