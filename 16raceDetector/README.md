 

# DataRace Rules

> Later, I realized that the code for each example should be placed in separate folders to avoid unwanted interactions. 

## 1goroutines

### Error

```golang
func Test_Race_goroutines(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)

	count := 0 // ----- race ----->
    
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count++ // <----- race ----- ( X many )
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
var mu sync.Mutex // correct (1/3) !

func Test_Race_goroutines(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)

	count := 0
    
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()   // correct (2/3) !
			count++
			mu.Unlock() // correct (3/3) !
		}()
	}
    
	wg.Wait()
}
```

### Enhanced

```go
func Test_Race_goroutines(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	var count int32 = 0
    
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // correct (1/1) !
		}()
	}
    
	wg.Wait()
}
```

### Operation

```bash
$ cd ./rules/1goroutine

$ make race
# go test -race -v -run Test_Race_goroutines ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/1goroutine/race 0.022s
# FAIL

# go test -race -v -run Test_Race_goroutines ./fixed/ | tail -n 3
# --- PASS: Test_Race_goroutines (0.00s)
# PASS
# ok      ./rules/1goroutine/fixed        0.028s

# go test -race -v -run Test_Race_goroutines ./enhanced/ | tail -n 3
# --- PASS: Test_Race_goroutines (0.00s)
# PASS
# ok      ./rules/1goroutine/enhanced      0.037s

$ make benchmark
# go test -v -bench=. -run=none -benchmem ./fixed/
# goos: linux
# goarch: amd64
# pkg: ./rules/1goroutine/fixed
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_goroutines
# Benchmark_Race_goroutines-8     70143081                16.87 ns/op            0 B/op          0 allocs/op
# PASS
# ok      ./rules/1goroutine/fixed        1.205s

# go test -v -bench=. -run=none -benchmem ./enhanced/
# goos: linux
# goarch: amd64
# pkg: ./rules/1goroutine/enhanced
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_goroutines
# Benchmark_Race_goroutines-8     161028873                7.057 ns/op           0 B/op   #        0 allocs/op
# PASS
# ok      ./rules/1goroutine/enhanced      1.894s
```

## 2sync

### Error

```go
func Test_Race_sync(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	count := 0 // ----- race ----->

	/*
		Two goroutines did not use the same mutex
		Disable the sync mechanism here will cause a race condition,
		which is my purpose of doing so
		(我故意的)
	*/
	mu := sync.Mutex{}
	mu2 := sync.Mutex{}
    
	go func() {
		defer wg.Done()
		mu.Lock()   // <- sync mechanism -
		count++     // <----- race ----- ( X 1 )
		mu.Unlock() // <- sync mechanism -
	}()
	go func() {
		defer wg.Done()
		mu2.Lock()   // <- sync mechanism -
		count++      // <----- race ----- ( X 1 )
		mu2.Unlock() // <- sync mechanism -
	}()
    
	wg.Wait()
}

```

### Fixed

```go
func Test_Race_sync(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	count := 0
    
	mu := sync.Mutex{}
	// mu2 := sync.Mutex{} // correct (1/3) !
    
	go func() {
		defer wg.Done()
		mu.Lock()
		count++
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()   // correct (2/3) !
		count++
		mu.Unlock() // correct (3/3) !
	}()
    
	wg.Wait()
}
```

### Enhanced

````go
// (empty)
````

### Operation

``` bash
$ cd ./rules/2sync

$ make race
# go test -race -v -run Test_Race_sync ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/2sync/race      0.017s
# FAIL

# go test -race -v -run Test_Race_sync ./fixed/ | tail -n 3
# --- PASS: Test_Race_sync (0.00s)
# PASS
# ok      ./rules/2sync/fixed     0.029s
```

## 3sync2

### Error

```go
func Test_Race_sync2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	var count int // ----- race ----->
    
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		mu.Lock()
		go func() {
			defer wg.Done()
			mu.Unlock() // - Unlock mutex early ! -
			count++     // <----- race -----
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
func Test_Race_sync2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	var count int
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		mu.Lock()
		go func() {
			defer wg.Done()
			count++
			mu.Unlock() // Correct (1/1) !
		}()
	}
    
	wg.Wait()
}
```

### Enhanced

````go
// (empty)
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_sync2 ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/3sync2/race     0.023s
# FAIL

# go test -race -v -run Test_Race_sync2 ./fixed/ | tail -n 3
# --- PASS: Test_Race_sync2 (0.01s)
# PASS
# ok      ./rules/3sync2/fixed    (cached)
```

## 4 map

### Error

```go
func Test_Race_map(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	var count = make(map[string]int) // ----- race ----->
    
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count["key"]++ // <----- race ----- ( X many )
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
var mu sync.Mutex // correct (1/3) !

func Test_Race_map(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    var count = make(map[string]int)
    
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()      // correct (2/3) !
			count["key"]++
			mu.Unlock()    // correct (3/3) !
		}()
	}
    
	wg.Wait()
}
```

### Enhanced

````go
func Test_Race_map(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	var count = sync.Map{}         // correct (1/4) !
	count.Store("key", new(int64)) // correct (2/4) !
    
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			value, _ := count.Load("key")      // correct (3/4) !
			atomic.AddInt64(value.(*int64), 1) // correct (4/4) !
		}()
	}
    
	wg.Wait()
    
	var sum int64
	count.Range(func(key, value interface{}) bool {
		sum = sum + *value.(*int64)
		return true
	})
	fmt.Println("result of sync map is ", sum)
}
````

### Operation

``` bash
$ go test -v -run Test_Race_map ./enhanced/
# === RUN   Test_Race_map
# result of sync map is  1000
# --- PASS: Test_Race_map (0.00s)
# PASS
# ok      ./rules/4map/enhanced

$ make race
# go test -race -v -run Test_Race_map ./race/ | tail -n 3
#        ./rules/4map/race/map_race_test.go:19 +0x78
# FAIL    ./rules/4map/race       0.020s
# FAIL

# go test -race -v -run Test_Race_map ./fixed/ | tail -n 3
# --- PASS: Test_Race_map (0.00s)
# PASS
# ok      ./rules/4map/fixed      (cached)

# go test -race -v -run Test_Race_map ./enhanced/ | tail -n 3
# --- PASS: Test_Race_map (0.00s)
# PASS
# ok      ./rules/4map/enhanced   (cached)

$ go test -v -bench=. -run=none -benchmem ./fixed/
# goos: linux
# goarch: amd64
# pkg: github.com/panhongrainbow/golangExplainedz/16raceDetector/rules/4map/fixed
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_map
# Benchmark_Race_map-8     2300634               576.1 ns/op            19 B/op          1 allocs/op
# PASS
# ok      ./rules/4map/fixed      1.869s

# go test -v -bench=. -run=none -benchmem ./enhanced/
# goos: linux
# goarch: amd64
# pkg: github.com/panhongrainbow/golangExplainedz/16raceDetector/rules/4map/enhanced
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_map
# Benchmark_Race_map-8     3202636               399.9 ns/op            17 B/op          1 allocs/op
# PASS
# ok      ./rules/4map/enhanced   1.667s
```

## 5channel

### Error

```go
func Test_Race_channel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	count := 0 // ----- race ----->

	
	c := make(chan int)
	close(c) // <- race -

	// First mover
	go func() {
		defer wg.Done()
		count++ // <----- race ----- ( X 1 )
		c <- 1  // <- channel mechanism -
	}()

	// Second mover
	go func() {
		defer wg.Done()
		<-c     // <- channel mechanism -
		count++ // <----- race ----- ( X 1 )
	}()
    
	wg.Wait()
}
```

### Fixed

```go
func Test_Race_channel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	count := 0
	c := make(chan int)

	/*
		Fix the race condition by maintaining the normal operation of the channel
	*/
	// Close(c) // correct (1/1)
    
	go func() {
		defer wg.Done()
		count++
		c <- 1
	}()
    
	go func() {
		defer wg.Done()
		<-c
		count++
	}()
    
	wg.Wait()
}
```

### Enhanced

````go
// (empty)
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_channel ./race/ | tail -n 3
# ./rules/5channel/race/channel_race_test.go:29 +0x185
# FAIL    ./rules/5channel/race   0.019s
# FAIL

# go test -race -v -run Test_Race_channel ./fixed/ | tail -n 3
# --- PASS: Test_Race_channel (0.00s)
# PASS
# ok      ./rules/5channel/fixed  0.029s
```

## 6channel2

### Error

```go
func Test_Race_channel2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
    
	c := make(chan int) // ----- race ----->
	defer close(c)
    
	for i := 0; i < 10; i++ {
		// i := i // <- race - // race because comment this line
		go func() {
			defer wg.Done()
			fmt.Println("send", i)
			c <- i // <----- race -----
		}()
	}
    
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			v := <-c
			fmt.Println("receive", v)
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
func Test_Race_channel2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(20)
    
	c := make(chan int)
	defer close(c)
    
	for i := 0; i < 10; i++ {
		i := i // Correct (1/1) !
		go func() {
			defer wg.Done()
			fmt.Println("send", i)
			c <- i
		}()
	}
    
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			v := <-c
			fmt.Println("receive", v)
		}()
	}
    
	wg.Wait()
}

```

### Enhanced

````go
// (empty)
````

### Operation

``` bash
$ make run
# go test -v -run Test_Race_channel2 ./race/
# === RUN   Test_Race_channel2
# send 10
# receive 10
# send 10
# send 10
# receive 10
# receive 10
# send 10
# receive 10

$ make race
# go test -race -v -run Test_Race_channel2 ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/6channel2/race  0.015s
# FAIL
# go test -race -v -run Test_Race_channel2 ./fixed/ | tail -n 3
# --- PASS: Test_Race_channel2 (0.00s)
# PASS
# ok      ./rules/6channel2/fixed (cached)
```

## 7timer

### Error

```go
func Test_Race_timer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	var count int // ----- race ----->
    
	go func() {
		defer wg.Done()
		count++ // <----- race -----
	}()
    
	select {
	case <-time.After(1 * time.Second): // <- random race -
		go func() {
			defer wg.Done()
			count++ // <----- race -----
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
var mu sync.Mutex // add (1/5) !

func Test_Race_timer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	var count int // ----- race ----->
    
	go func() {
		defer wg.Done()
		mu.Lock()   // add (2/5) !
		count++     // <----- race -----
		mu.Unlock() // add (3/5) !
	}()
    
	select {
	case <-time.After(1 * time.Second): // <- random race -
		go func() {
			defer wg.Done()
			mu.Lock()   // add (4/5) !
			count++     // <----- race -----
			mu.Unlock() // add (5/5) !
		}()
	}
    
	wg.Wait()
}
```

### Enhanced

````go
func Test_Race_timer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	var count int32 // ----- race ----->
    
	go func() {
		defer wg.Done()
		atomic.AddInt32(&count, 1) // correct (1/2) !
	}()
    
	select {
	case <-time.After(1 * time.Second): // <- random race -
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // correct (2/2) !
		}()
	}
    
	wg.Wait()
}
````

### Operation

``` bash
$ make race

# go test -race -v -run Test_Race_timer ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/7timer/race     1.019s
# FAIL

# go test -race -v -run Test_Race_timer ./fixed/ | tail -n 3
# --- PASS: Test_Race_timer (1.00s)
# PASS
# ok      ./rules/7timer/fixed    (cached)

# go test -race -v -run Test_Race_timer ./enhanced/ | tail -n 3
# --- PASS: Test_Race_timer (1.00s)
# PASS
# ok      ./rules/7timer/enhanced 1.034s
```

## 8timer2

### Error

```go
func Test_Race_timer2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
    
	var count int // ----- race ----->
    
	go func() {
		defer wg.Done()
		count++ // <----- race -----
	}()
    
	go func() {
		defer wg.Done()
		select {
		case <-timer.C: // <- random race -
			// if timer is setted to 0 * time.Second
		default:
			// if timer is setted to 1 * time.Second
			count++ // <----- race -----
		}
	}()
    
	wg.Wait()
}
```

### Fixed

```go
var mu sync.Mutex // add (1/5) !

func Test_Race_timer2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	var count int
    
	go func() {
		defer wg.Done()
		mu.Lock()   // add (2/5) !
		count++
		mu.Unlock() // add (3/5) !
	}()
    
	go func() {
		defer wg.Done()
		select {
		case <-timer.C:
			// if timer is setted to 0 * time.Second
		default:
			// if timer is setted to 1 * time.Second
			mu.Lock()   // add (4/5) !
			count++
			mu.Unlock() // add (5/5) !
		}
	}()
    
	wg.Wait()
}
```

### Enhanced

````go
func Test_Race_timer2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
    
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
    
	var count int32
    
	go func() {
		defer wg.Done()
		atomic.AddInt32(&count, 1) // correct (1/2) !
	}()
    
	go func() {
		defer wg.Done()
		select {
		case <-timer.C:
			// if timer is setted to 0 * time.Second
		default:
			// if timer is setted to 1 * time.Second
			atomic.AddInt32(&count, 1) // correct (1/2) !
		}
	}()
    
	wg.Wait()
}
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_timer2 ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/8timer2/race    0.015s
# FAIL

# go test -race -v -run Test_Race_timer2 ./fixed/ | tail -n 3
# --- PASS: Test_Race_timer2 (0.00s)
# PASS
# ok      ./rules/8timer2/fixed   0.026s

# go test -race -v -run Test_Race_timer2 ./enhanced/ | tail -n 3
# --- PASS: Test_Race_timer2 (0.00s)
# PASS
# ok      ./rules/8timer2/enhanced        0.025s
```

## 9select

### Error

```go
var count int // ----- race ----->

func inc(ch1, ch2 chan bool) {
	select {
	case <-ch1:
	case <-ch2:
	}
	count++ // <----- race ----- ( X many )
}

func Test_Race_select(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
    
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() { // <- race -
			defer wg.Done()
			inc(ch1, ch2) // <- race -
		}()
	}

	// What if close channels // <- race -
	close(ch1)
	close(ch2)
    
	wg.Wait()
}
```

### Fixed

```go
var mu sync.Mutex // add (1/5) !

var count int

func inc(ch1, ch2 chan bool) {
	mu.Lock()   // add (2/5) !
	count++
	mu.Unlock() // add (3/5) !
	select {
	case <-ch1:
	case <-ch2:
	}
	mu.Lock()   // add (4/5) !
	count++
	mu.Unlock() // add (5/5) !
}

func Test_Race_select(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
    
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			inc(ch1, ch2)
		}()
	}
    
	close(ch1)
	close(ch2)
    
	wg.Wait()
}
```

### Enhanced

````go
var count int32

func inc(ch1, ch2 chan bool) {
	atomic.AddInt32(&count, 1) // correct (1/2) !
	select {
	case <-ch1:
	case <-ch2:
	}
	atomic.AddInt32(&count, 1) // correct (2/2) !
}

func Test_Race_select(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
    
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			inc(ch1, ch2)
		}()
	}
    
	close(ch1)
	close(ch2)
    
	wg.Wait()
}
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_select ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/9select/race    0.029s
# FAIL
# go test -race -v -run Test_Race_select ./fixed/ | tail -n 3
# --- PASS: Test_Race_select (0.00s)
# PASS
# ok      ./rules/9select/fixed   0.027s
# go test -race -v -run Test_Race_select ./enhanced/ | tail -n 3
# --- PASS: Test_Race_select (0.00s)
# PASS
# ok      ./rules/9select/enhanced        0.031s
```

## 10interface

### Error

```go
type I interface {
	Set(x int)
}

type T struct {
	x int
}

func (t *T) Set(x int) {
	t.x = x
}

var obj I

func Write() {
	obj = new(T)
}

func Set() {
	obj = new(T)
	obj.Set(10)
}

func Test_Race_interface(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			Write() // <- race -
		}()

		go func() {
			defer wg.Done()
			Set() // <- race -
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
type I interface {
	Set(x int)
}

type T struct {
	x int
}

func (t *T) Set(x int) {
	t.x = x
}

var obj I

func Write() {
	obj = new(T)
}

func Set() {
	obj = new(T)
	obj.Set(10)
}

func MutexWrite() {
	mu.Lock()    // add (1/4) !
	obj = new(T)
	mu.Unlock()  // add (2/4) !
}

func MutexSet() {
	mu.Lock() // add (3/4) !
	obj = new(T)
	obj.Set(10)
	mu.Unlock() // add (4/4) !
}

func Test_fixed_interface(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			MutexWrite() // fixed (1/2) !
		}()

		go func() {
			defer wg.Done()
			MutexSet() // fixed (2/2) !
		}()
	}
    
	wg.Wait()
}
```

### Enhanced

````go
type I interface {
	Set(x int)
}

type T struct {
	x int
}

func (t *T) Set(x int) {
	t.x = x
}

var obj I

func Write() {
	obj = new(T)
}

func Set() {
	obj = new(T)
	obj.Set(10)
}

var mu2 sync.Mutex

func AtomicWrite() {
	// obj2.Store(new(T)) // Big no, no, not atomic !!!
	t := new(T)
	obj2.Store(t) // Atomic !!! // correct (1/2) !
}

func AtomicSet() {
	mu2.Lock()
	obj2.Load().(*T).Set(10) // Big no, no, not atomic !!! need to use mutex // correct (2/2) !
	mu2.Unlock()
}
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_interface ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/10interface/race        0.032s
# FAIL
# go test -race -v -run Test_fixed_interface ./race/ | tail -n 3
# --- PASS: Test_fixed_interface (0.00s)
# PASS
# ok      ./rules/10interface/race        0.040s
# go test -race -v -run Test_atomic_interface ./race/
# === RUN   Test_atomic_interface
# --- PASS: Test_atomic_interface (0.00s)
# PASS
# ok      ./rules/10interface/race        0.029s

$ go test -v -bench=. -run=none -benchmem ./race/
# goos: linux
# goarch: amd64
# pkg: ./rules/10interface/race
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_fixed_interface
# Benchmark_Race_fixed_interface-8         2049138               624.9 ns/op            19 B/op          2 allocs/op
# Benchmark_Race_atomic_interface
# Benchmark_Race_atomic_interface-8        2428218               500.5 ns/op             8 B/op          1 allocs/op
# PASS
# ok      ./rules/10interface/race        3.608s
```

## 11closure

### Error

```go
import (
	"sync"
	"testing"
)

func Test_Race_closure(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	count := 0 // ----- race ----->
    
	closure := func() func() {
		var mu sync.Mutex // <- race -
		return func() {
			defer wg.Done()
			mu.Lock()
            count++ // <----- race ----- ( X many )
			mu.Unlock()
		}
	}

	/*for i := 0; i < 1000; i++ {
		go closure() //  Not correct, it will result in being unable to unlock
	}*/

	// Create 1000 goroutines
	for i := 0; i < 1000; i++ {
		/*
			The return value should be stored in the fn variable to be garbage collected;
			Otherwise, it will result in being unable to unlock
		*/
		fn := closure()
		go func() {
			fn() // Call the fn variable instead of closure()
		}()
	}
    
	wg.Wait()
}
```

### Fixed

```go
func Test_Race_closure(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	count := 0
    
	var mu sync.Mutex // correct (1/1) !
    
	closure := func() func() {
		return func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}
	}
    
	for i := 0; i < 1000; i++ {
		fn := closure()
		go func() {
			fn()
		}()
	}
    
	wg.Wait()
}

```

### Enhanced

````go
import (
	"sync"
	"sync/atomic"
	"testing"
)

func Test_Race_closure(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	var count int32 // correct (1/2) !
    
	closure := func() func() {
		return func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // correct (2/2) !
		}
	}
    
	for i := 0; i < 1000; i++ {
		fn := closure()
		go func() {
			fn()
		}()
	}
    
	wg.Wait()
}
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_closure ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/11closure/race  0.024s
# FAIL

# go test -race -v -run Test_Race_closure ./fixed/ | tail -n 3
# --- PASS: Test_Race_closure (0.00s)
# PASS
# ok      ./rules/11closure/fixed (cached)

# go test -race -v -run Test_Race_closure ./enhanced/ #| tail -n 3
# === RUN   Test_Race_closure
# --- PASS: Test_Race_closure (0.00s)
# PASS
# ok      ./rules/11closure/enhanced      (cached)
```

## 12list

### Error

```go
type List struct {
	value int
	next  *List
}

func Test_Race_list(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	root := &List{value: -1} // ----- race ----->
    
	for i := 0; i < 1000; i++ { // <- race -
		i := i
		go func() {
			defer wg.Done()
			list := &List{value: i}
			next := root
			for {
				if next.next == nil {
					next.next = list // <----- race ----- ( X many )
					break
				} else {
					next = next.next
				}
			}
		}()
	}
	wg.Wait()
    
	var count int
	next := root
	for {
		if next.next == nil {
			break
		} else {
			count++
			next = next.next
		}
	}
	fmt.Println("list length: ", count)
}
```

### Fixed

```go
var mu sync.Mutex // add (1/3) !

type List struct {
	value int
	next  *List
}

func Test_Race_list(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	root := &List{value: -1}
    
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock() // add (2/3) !
			list := &List{value: i}
			next := root
			for {
				if next.next == nil {
					next.next = list
					break
				} else {
					next = next.next
				}
			}
			mu.Unlock() // add (3/3) !
		}()
	}
    
	wg.Wait()
    
	var count int
	next := root
	for {
		if next.next == nil {
			break
		} else {
			count++
			next = next.next
		}
	}
	fmt.Println("list length: ", count)
}
```

### Enhanced

````go
type List struct {
	value int
	next  atomic.Pointer[List] // correct (1/3) !
}

func Test_Race_list(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
    
	root := &List{value: -1}
    
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			defer wg.Done()
			list := List{value: i}
			next := root
			for {
				if next.next.Load() == nil { // correct (2/3) !
					if next.next.CompareAndSwap(nil, &list) { // correct (3/3) !
						break
					}
				} else {
					next = next.next.Load()
				}
			}
		}()
	}
    
	wg.Wait()
    
	var count int
	next := root
	for {
		if next.next.Load() == nil {
			break
		} else {
			count++
			next = next.next.Load()
		}
	}
	fmt.Println("list length: ", count)
    
	fmt.Println(root.value)
	fmt.Println(root.next.Load().value)
	fmt.Println(root.next.Load().next.Load().value)
	fmt.Println(root.next.Load().next.Load().next.Load().value)
	fmt.Println(root.next.Load().next.Load().next.Load().next.Load().value)
}
````

### Operation

``` bash
$ make race
# go test -race -v -run Test_Race_list ./race/ | tail -n 3
# FAIL
# FAIL    ./rules/12list/race     0.034s
# FAIL

# go test -race -v -run Test_Race_list ./fixed/ | tail -n 3
# --- PASS: Test_Race_list (0.03s)
# PASS
# ok      ./rules/12list/fixed    0.063s

# go test -race -v -run Test_Race_list ./enhanced/ | tail -n 3
# --- PASS: Test_Race_list (0.03s)
# PASS
# ok      ./rules/12list/enhanced 0.056s

$ go test -v -bench=. -run=none -benchmem ./fixed/
# goos: linux
# goarch: amd64
# pkg: ./rules/12list/fixed
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_list
# Benchmark_Race_list-8                865           1446464 ns/op           48707 B/op       2007 allocs/op
# PASS
# ok      ./rules/12list/fixed    1.397s

# go test -v -bench=. -run=none -benchmem ./enhanced/
# goos: linux
# goarch: amd64
# pkg: ./rules/12list/enhanced
# cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
# Benchmark_Race_list
# Benchmark_Race_list-8               2365            554933 ns/op           48041 B/op       2002 allocs/op
# PASS
# ok      ./rules/12list/enhanced 1.369s
```

## 13

### Error

```go
1
```

### Fixed

```go
2
```

### Enhanced

````go
3
````

### Operation

``` bash
4
```













































































