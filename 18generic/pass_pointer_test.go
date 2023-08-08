package generic

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// >>>>> >>>>> >>>>> Canine

type Canine interface {
	Bark()
	Feed(int)
	Scale() int
	Clone() any
}

// >>>>> >>>>> >>>>> Bulldog

type Bulldog struct{ Weight int }

func (*Bulldog) Bark() {
	fmt.Println("Woof from Bulldog !")
}

func (bulldog *Bulldog) Feed(weight int) {
	bulldog.Weight = bulldog.Weight + weight
}

func (bulldog *Bulldog) Scale() int {
	return bulldog.Weight
}

func (bulldog *Bulldog) Clone() any {
	copyDog := &Bulldog{}
	return copyDog
}

// >>>>> >>>>> >>>>> Corgi

type Corgi struct{ Weight int }

func (*Corgi) Bark() {
	fmt.Println("Woof from Corgi !")
}

func (corgi *Corgi) Feed(weight int) {
	corgi.Weight = corgi.Weight + weight
}

func (corgi *Corgi) Scale() int {
	return corgi.Weight
}

func (corgi *Corgi) Clone() any {
	copyDog := &Corgi{}
	return copyDog
}

// >>>>> >>>>> >>>>> play

func PlayWithMirror[dogT Canine](dog dogT) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	dog.Bark()
	dog.Feed(1)
	var mirror dogT
	mirror.Bark()
	mirror.Feed(1)
}

func Birth[T any, dogT Canine](dog dogT) T {
	dog.Feed(-3)
	return dog.Clone().(T)
}

func HasPregnant(dog Canine) any {
	dog.Feed(-3)
	return dog.Clone()
}

func BulldogHasPregnant(dog *Bulldog) *Bulldog {
	dog.Feed(-3)
	return dog.Clone().(*Bulldog)
}

func CorgiHasPregnant(dog *Corgi) *Corgi {
	dog.Feed(-3)
	return dog.Clone().(*Corgi)
}

// >>>>> >>>>> >>>>> run

func Test_Check_PassPointer(t *testing.T) {
	bulldog := new(Bulldog)
	bulldog.Feed(5)
	require.Equal(t, 5, bulldog.Scale())

	PlayWithMirror(bulldog)
	PlayWithMirror(&Corgi{})
	require.Equal(t, 6, bulldog.Scale())

	puppy := Birth[*Bulldog](bulldog)
	require.Equal(t, 3, bulldog.Scale())
	puppy.Bark()
}

func Benchmark_Performance_NoInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bulldog := new(Bulldog)
		bulldogPuppy := BulldogHasPregnant(bulldog)
		bulldogPuppy.Feed(1)

		corgi := new(Corgi)
		corgiPuppy := CorgiHasPregnant(corgi)
		corgiPuppy.Feed(1)
	}
}

func Benchmark_Performance_NoGeneric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bulldog := new(Bulldog)
		bulldogPuppy := HasPregnant(bulldog)
		bulldogPuppy.(*Bulldog).Feed(1)

		corgi := new(Corgi)
		corgiPuppy := HasPregnant(corgi)
		corgiPuppy.(*Corgi).Feed(1)
	}
}

func Benchmark_Performance_PassPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bulldog := new(Bulldog)
		bulldogPuppy := Birth[*Bulldog](bulldog)
		bulldogPuppy.Feed(1)

		corgi := new(Corgi)
		corgiPuppy := Birth[*Corgi](corgi)
		corgiPuppy.Feed(1)
	}
}

func Benchmark_Performance_PassInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bulldog Canine = new(Bulldog)
		bulldogPuppy := Birth[Canine](bulldog)
		bulldogPuppy.Feed(1)

		var corgi Canine = new(Corgi)
		corgiPuppy := Birth[Canine](corgi)
		corgiPuppy.Feed(1)
	}
}
