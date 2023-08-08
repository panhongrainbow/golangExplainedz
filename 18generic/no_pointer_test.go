package generic

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// >>>>> >>>>> >>>>> Feline

type Feline interface {
	Meow()
	Feed(int)
	Scale() int
	Clone() any
}

// >>>>> >>>>> >>>>> Calico

type Calico struct{ Weight int }

func (Calico) Meow() {
	fmt.Println("Meow from Calico !")
}

func (calico Calico) Feed(weight int) {
	calico.Weight = calico.Weight + weight
}

func (calico Calico) Scale() int {
	return calico.Weight
}

func (calico Calico) Clone() any {
	copyCat := Calico{}
	return copyCat
}

// >>>>> >>>>> >>>>> Tuxedo

type Tuxedo struct{ Weight int }

func (Tuxedo) Meow() {
	fmt.Println("Meow from Tuxedo !")
}

func (tuxedo Tuxedo) Feed(weight int) {
	tuxedo.Weight = tuxedo.Weight + weight
}

func (tuxedo Tuxedo) Scale() int {
	return tuxedo.Weight
}

func (tuxedo Tuxedo) Clone() any {
	copyDog := Tuxedo{}
	return copyDog
}

// >>>>> >>>>> >>>>> play

func PuffUpWithMirror[catT Feline](cat catT) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	cat.Meow()
	cat.Feed(1)
	var mirror catT
	mirror.Meow()
	mirror.Feed(1)
}

func Produce[T any, catT Feline](cat catT) T {
	cat.Feed(-3)
	return cat.Clone().(T)
}

// >>>>> >>>>> >>>>> run

func Test_Check_NoPointer(t *testing.T) {
	calico := Calico{}
	calico.Feed(5)
	require.Equal(t, 0, calico.Scale())

	PuffUpWithMirror(calico)
	PuffUpWithMirror(Tuxedo{})
	require.Equal(t, 0, calico.Scale())

	kitten := Produce[Calico](calico)
	require.Equal(t, 0, kitten.Scale())
	kitten.Meow()
}

func Benchmark_Performance_NoPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calico := new(Calico)
		calicoKitten := Produce[Calico](calico)
		calicoKitten.Feed(1)

		tuxedo := new(Tuxedo)
		tuxedoKitten := Produce[Tuxedo](tuxedo)
		tuxedoKitten.Feed(1)
	}
}
