# Generic

## GC Shape

Because when using generics, the compiled program becomes larger due to expansion.

Additionally, considering the performance of the garbage collector, `generics are linked to GC shape` instead of variable types.

This concept is referred to as `monomorphization`.



However, monomorphization can lead to decreased program performance.

When the GC shape is `a primitive type or a struct, or when dealing with slices`, monomorphization usually proceeds smoothly without the need to extract iface from a dictionary.

(单态化会导致程序性能下降，GCShape 在处理原始类型或结构体、切片时，通常会顺利进行单态化，而不会涉及从字典获取接口（iface）的行为)



However, this behavior occurs when dealing with pointers. (指标有这问题)

This is why generics can result in slower code.

## Three Scenarios

The slowdown of generics can be divided into `the following three scenarios` for discussion.

When generics involve the following conditions (Three Reasons), the speed will become slower.

| Reason    | Description                                                  |
| --------- | ------------------------------------------------------------ |
| common    | As long as calling methods of the argument within the function.<br /><br />The main reason for `performance drops` usually stems from having `type parameters` in generic functions `declared as [T constraint]`, and then calling methods of T within those functions.<br />会形成 (泛型 interface) 2 (方法 interface) 转换和查找 |
| interface | When calling methods of the parameter within the function, but `passing an interface as the argument`, it becomes slower.<br /><br />Using `an interface argument` results in `a conversion from one interface to another to access the method address of T`.<br />This is what causes the slowness.<br /><br />Moreover, `the resulting performance degradation` is most significant.<br />(Conversion from one interface to another)<br /><br />这个最慢，因为会形成 (泛型 interface) 2 (方法 interface) 的查找，找最久 |
| pointer   | Pointers can stress the garbage collector.<br /><br />The GC needs to verify if `the pointer` can be collected and then check if `the data it points to` can also be collected.<br />This leads to `a double-checking process`.<br />The issue of `GC pressure caused by pointers` also exists when using interfaces.<br /><br />这是不只是在泛型会有的问题，在 interface 也会有这问题 |

## Worst Scenarios

The benchmark testing shows particularly slow results for the following cases.

```bash
$ cd /home/panhong/go/src/github.com/panhongrainbow/golangExplainedz/18generic

$ make group
```

The results are as follows

```go
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1 point 60 GHz
Benchmark_Group_NoGenericA-8          	477054699	         2.392 ns/op
Benchmark_Group_NoGenericB-8          	516782533	         2.295 ns/op
Benchmark_Group_InterfaceA-8          	767664978	         1.750 ns/op
Benchmark_Group_InterfaceB-8          	679029472	         1.575 ns/op
Benchmark_Group_GenericA-8            	420106927	         2.760 ns/op
Benchmark_Group_GenericB-8            	458548924	         2.641 ns/op
Benchmark_Group_GenericInterfaceA-8   	213647784	         5.676 ns/op # <<< worst
Benchmark_Group_GenericInterfaceB-8   	239801074	         4.934 ns/op # <<< worst
Benchmark_Group_GenericInterfaceC-8   	196562356	         6.047 ns/op # <<< worst
Benchmark_Group_DoAddNoGeneric-8      	561499844	         2.161 ns/op
PASS
```

### GenericInterfaceA

```go
func Benchmark_Group_GenericInterfaceA(b *testing.B) {
	var obj Adder = &A{} // Addr is an interface, and the interface is generated first
	for i := 0; i < b.N; i++ {
		DoAdd(obj) // Passing parameters into generics
	}
}
```

### GenericInterfaceB

```go
func Benchmark_Group_GenericInterfaceB(b *testing.B) {
	var obj Adder = &B{} // Addr is an interface, and the interface is generated first
	for i := 0; i < b.N; i++ {
		DoAdd(obj) // Passing parameters into generics
	}
}
```

### GenericInterfaceC

```go
func Benchmark_Group_GenericInterfaceC(b *testing.B) {
	var obj Adder = C{} // Addr is an interface, and the interface is generated first
	for i := 0; i < b.N; i++ {
		DoAdd(obj) // Passing parameters into generics
	}
}
```

## Performance hierarchy

### Five Layers

Perform benchmark tests to analyze performance under different conditions.

There are five layers

1. **Ultra-High Performance:** 极高性能层
2. **High Performance:** 高性能层
3. **Balanced Performance:** 均衡性能层
4. **Power Efficiency:** 能效性能层
5. **The worst Performance:** 不良性能层

```bash
$ cd /home/panhong/go/src/github.com/panhongrainbow/golangExplainedz/18generic

$ make benchmark
```

The results are as follows

```go
Benchmark_Performance_NoPointer-8       	25938903	        42.52 ns/op
Benchmark_Performance_NoInterface-8     	654768386	         1.898 ns/op
Benchmark_Performance_NoGeneric-8       	18926174	        68.53 ns/op
Benchmark_Performance_PassPointer-8     	14744830	        68.26 ns/op
Benchmark_Performance_PassInterface-8   	13418542	        87.64 ns/op
```

#### Ultra-High Performance

Cannot find any interfaces or any here !

```go
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

func BulldogHasPregnant(dog *Bulldog) *Bulldog {
	dog.Feed(-3)
	return dog.Clone().(*Bulldog)
}

func CorgiHasPregnant(dog *Corgi) *Corgi {
	dog.Feed(-3)
	return dog.Clone().(*Corgi)
}
```

Output：

```go
Benchmark_Performance_NoInterface-8     	654768386	         1.898 ns/op
```

#### High Performance

Without pointers, some values cannot be modified. This method simply cannot be used.

(`这例子不能使用，不能用指标，值都不能改`)



Here, the code uses generics, but `the passed parameter is not a pointer`.

(不使用指标当参数传入泛型函式)



Because pointers have an impact on the performance of the garbage collector GC, as they require the GC to `not only check if the pointers themselves can be collected, but also to inspect the data behind the pointers for collectivity`.

(GC 检查指标，也检查后面资料，共2次检查)



This essentially makes the GC perform two checks, so it's `not recommended to pass pointers as arguments to generic functions`.

(因为要2次检查，所以不建议在范型上使用)



The double-checking behavior of the GC due to pointers isn't exclusive to generics;

it also occurs when `assigning pointers to interfaces`.

However, `why is it particularly advised against using pointer parameters in generics?`

(GC检查2次的问题也会发生在接口，为何要特别限制泛型)



It might be related to the interaction between generics and `GC shapes`.

To improve GC efficiency, `variables in similar memory configurations share the same GC shape`.

This has already added complexity to memory management and is more complex than interfaces.

Given this increased complexity, it's `less recommended to use pointers as parameters in generic functions`.

(因为泛型和 GC Shape 有关，内存更复杂了，不要再用指标了)



As shown below, the code utilizes：

```go
calicoKitten := Produce[Calico](calico)

// AND

tuxedoKitten := Produce[Tuxedo](tuxedo)
```

Not as bellow：

```go
calicoKitten := Produce[*Calico](calico)

// AND

tuxedoKitten := Produce[*Tuxedo](tuxedo)
```

The complete code is as follows

```go
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
```

Output：

```go
Benchmark_Performance_NoPointer-8       	25938903	        42.52 ns/op
```

#### Balanced Performance

As follows, the function HasPregnant takes `a parameter of the Canine interface` and `returns any`.

The entire function uses `interfaces`.

(这里只使用接口)



The question is: ` when using pointers as parameters, the performance of interfaces and generics is similar`, so why is there a specific restriction on using pointers with generics?

This is due to the relationship between generics and GC shapes. It's already becoming complex and `can't become any more complicated`.

(为何要特别限制泛型使用指标？泛型和 GC Shape 有关，不能在复杂)

```go
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

func HasPregnant(dog Canine) any { // <<<<< use interfaces
	dog.Feed(-3)
	return dog.Clone()
}
```

Output：

```go
Benchmark_Performance_NoGeneric-8       	18926174	        68.53 ns/op
```

#### Power Efficiency

Here is where `passing pointer parameters to generic functions` is not recommended, but the execution result is quite `similar to the previous example that only uses interfaces`.

(不建议，但和只使用接口的版本效能接近，因为会有 iface 查找的问题)

```go
func Benchmark_Performance_PassPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bulldog := new(Bulldog)
		bulldogPuppy := Birth[*Bulldog](bulldog) // pass point *Bulldog
		bulldogPuppy.Feed(1)

		corgi := new(Corgi)
		corgiPuppy := Birth[*Corgi](corgi) // pass point *Corgi
		corgiPuppy.Feed(1)
	}
}
```

Output：

```go
Benchmark_Performance_PassPointer-8     	14744830	        68.26 ns/op
```

#### The Worst Performance

The following is when `an interface is passed as a parameter` to a generic, causing a lookup of both the generic interface and the method interface.

This is `the worst situation` to occur.

(造成 会形成 (泛型 interface) 2 (方法 interface) 的查找，惨)

``` go
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
```

Output：

```go
Benchmark_Performance_PassInterface-8   	13418542	        87.64 ns/op
```

## Get Field

### Explanation

Basically, within the function, it's `not` possible to directly `access the A and B of that tuple struct`.

Only through the interface method, this situation of `calling the function` can be avoided.

This is what is described as `common reason` in the three reasons for slowness.

(constrait 为 any 不指定，tuple struct 中的 A B Field 取不到的，又要 call function，为慢速中的 common 理由)



It should be said that the code within the generic function can only recognize the parameter as a constrained type, and the code `cannot make any assumptions about the type of the parameter beyond the constraints`.

(不能對一個參數的 type 做 constraint 外的任何假設)



In the case of `func Func(a Tuple[int,string])`, it's already instantiated, so accessing `a.A` and `a.B` is fine.

(无法取值)



However, in the case of `func Func[A any, B any, T Tuple[A,B]](a T)`, you `can't directly access `a.A` and `a.B` unless `T` is an interface that has methods to access them.

(可以取值)



The idea is to understand the topic without stepping on the slow-call method pitfall.

(这里会中 common reason 的快速雷)



But the current syntax `forces you into this pitfall`, unless you bring variables in and then don't access its contents.

(除非函式把函式拿进来后不存取，不可能)



The problem is, if you do that, the generic function becomes useless, and for tasks like maps or sorts, `you still need to care about the item value`.

(不可能，map 和 sort 都需要取值，不可能 拿进来后不存取)

### Example

In the following example, the generic constraint is set to `any`, without specifying a type, which will lead to `an inability to retrieve values`.

(泛型没有指定型态，会无法取值)

```go
package generic

import (
	"fmt"
	"testing"
)

type Tuple struct {
	FieldA int
	FieldB string
}

func GenericFunction[T any](param T) { // No specific type is specified for generics.
	switch v := param.(type) { // Here's an error, unable to retrieve values
	case Tuple:
		fmt.Printf("Tuple FieldA: %d, FieldB: %s\n", v.FieldA, v.FieldB)
	default:
		fmt.Println("Unsupported type")
	}
}

func Test_GetField(t *testing.T) {
	tupleInstance := Tuple{FieldA: 42, FieldB: "Hello"}

	GenericFunction(tupleInstance)
}
```

To solve the example above, you'll need to use a function call, as shown below

(就用 call function 去解上面的问题)

```go
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
```

Output：

```go
FieldA: 42, FieldB: Hello
```

There's another approach, which involves using a function to encapsulate `the type parameter T`.

(T type parameter 被函式 F 包装隔离了)



As shown below, the `Sort` function takes parameters `a []T` and `compare F`, where `a []T` is processed by the `compare F` function without direct involvement from the `Sort` function.

(函式 F 隔离 T，Sort 函式不必介入)

```go
package main

import (
	"fmt"
	"sort"
)

func Sort[T any, F func(a T, b T) bool](a []T, compare F) { // F encapsulates a
	sort.Slice(a, func(i, j int) bool { // i j is the index of slice.
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
```

Output：

```go
[1 2 3 5 8]
[1.4 2.2 3.5 5.1 8.3]
```

