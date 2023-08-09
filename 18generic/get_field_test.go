package generic

import (
	"fmt"
	"testing"
)

type FieldGetter interface {
	GetFieldA() int
	GetFieldB() string
}

type Tuple struct {
	FieldA int
	FieldB string
}

func (t Tuple) GetFieldA() int {
	return t.FieldA
}

func (t Tuple) GetFieldB() string {
	return t.FieldB
}

func GenericFunction[T FieldGetter](param T) {
	fieldA := param.GetFieldA()
	fieldB := param.GetFieldB()
	fmt.Printf("FieldA: %d, FieldB: %s\n", fieldA, fieldB)
}

func Test_Check_GetField(t *testing.T) {
	tupleInstance := Tuple{FieldA: 42, FieldB: "Hello"}

	GenericFunction(tupleInstance)
}
