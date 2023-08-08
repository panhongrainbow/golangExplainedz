package generic

import "testing"

// >>>>> >>>>> >>>>> Adder

type Adder interface {
	Add()
}

// >>>>> >>>>> >>>>> A

type A struct {
	num1 uint64
	num2 int64
}

func (a *A) Add() {
	a.num1++
	a.num2 = int64(a.num1 / 2)
}

// >>>>> >>>>> >>>>> B

type B struct {
	num1 uint64
	num2 int64
}

func (b *B) Add() {
	b.num1++
	b.num2 = int64(b.num1 / 2)
}

// >>>>> >>>>> >>>>> C

type C struct {
	num1 uint64
	num2 int64
}

func (c C) Add() {
	c.num1++
	c.num2 = int64(c.num1 / 2)
}

// >>>>> >>>>> >>>>> Method

func DoAdd[T Adder](t T) {
	t.Add()
}

func DoAddInterface(t Adder) {
	t.Add()
}

func DoAddNoGeneric(a Adder) {
	a.Add()
}

// >>>>> >>>>> >>>>> Benchmark

func Benchmark_Group_NoGenericA(b *testing.B) {
	obj := &A{}
	for i := 0; i < b.N; i++ {
		obj.Add()
	}
}

func Benchmark_Group_NoGenericB(b *testing.B) {
	obj := &B{}
	for i := 0; i < b.N; i++ {
		obj.Add()
	}
}

func Benchmark_Group_InterfaceA(b *testing.B) {
	obj := &A{}
	for i := 0; i < b.N; i++ {
		DoAddInterface(obj)

	}
}

func Benchmark_Group_InterfaceB(b *testing.B) {
	obj := &B{}
	for i := 0; i < b.N; i++ {
		DoAddInterface(obj)

	}
}

func Benchmark_Group_InterfaceC(b *testing.B) {
	obj := &C{}
	for i := 0; i < b.N; i++ {
		DoAddInterface(obj)

	}
}

func Benchmark_Group_GenericA(b *testing.B) {
	obj := &A{}
	for i := 0; i < b.N; i++ {
		DoAdd(obj)
	}
}

func Benchmark_Group_GenericB(b *testing.B) {
	obj := &B{}
	for i := 0; i < b.N; i++ {
		DoAdd(obj)
	}
}

func Benchmark_Group_GenericInterfaceA(b *testing.B) {
	var obj Adder = &A{}
	for i := 0; i < b.N; i++ {
		DoAdd(obj)
	}
}

func Benchmark_Group_GenericInterfaceB(b *testing.B) {
	var obj Adder = &B{}
	for i := 0; i < b.N; i++ {
		DoAdd(obj)
	}
}

func Benchmark_Group_GenericInterfaceC(b *testing.B) {
	var obj Adder = C{}
	for i := 0; i < b.N; i++ {
		DoAdd(obj)
	}
}

func Benchmark_Group_DoAddNoGeneric(b *testing.B) {
	var obj Adder = &A{}
	for i := 0; i < b.N; i++ {
		DoAddNoGeneric(obj)
	}
}
