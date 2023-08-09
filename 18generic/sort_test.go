package generic

import (
	"fmt"
	"sort"
	"testing"
)

func Sort[T any, F func(a T, b T) bool](a []T, compare F) {
	sort.Slice(a, func(i, j int) bool {
		return compare(a[i], a[j])
	})
}

func Test_Check_sort(t *testing.T) {
	intSlice := []int{5, 2, 8, 1, 3}

	Sort(intSlice, func(a, b int) bool {
		return a < b
	})

	fmt.Println(intSlice)

	floatSlice := []float32{5.1, 2.2, 8.3, 1.4, 3.5}

	Sort(floatSlice, func(a, b float32) bool {
		return a < b
	})

	fmt.Println(floatSlice)
}
