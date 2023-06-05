package examples

import (
	"errors"
	"sync"
	"testing"
)

var errorPool = sync.Pool{
	New: func() interface{} {
		return &MyError{}
	},
}

var errorPool2 = sync.Pool{
	New: func() interface{} {
		return errors.New("")
	},
}

var errorPool3 = sync.Pool{
	New: func() interface{} {
		var err error
		return err
	},
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func (e *MyError) ReSet() {
	e.Msg = ""
	return
}

func (e *MyError) Set(str string) {
	e.Msg = str
	return
}

/*func Test_err_join(t *testing.T) {
	for i := 0; i < 1000; i++ {
		err1 := errorPool.Get().(error)
		err1 = &MyError{Msg: "error 1"}

		err2 := errorPool.Get().(error)
		err2 = &MyError{Msg: "error 2"}

		err3 := errorPool.Get().(error)
		err3 = &MyError2{Msg: "error 3"}

		err5 := errors.Join(err1, err2, err3)

		var target *MyError
		var target2 *MyError2

		require.False(t, errors.As(err2, &target2))
		require.True(t, errors.As(err5, &target))
		require.True(t, errors.As(err5, &target2))

		err1.(*MyError).ReSet()
		errorPool.Put(err1)
		err2.(*MyError).ReSet()
		errorPool.Put(err2)
		err3.(*MyError2).ReSet()
		errorPool.Put(err3)
	}
}*/

func Benchmark_err_join(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err1 := errorPool.Get().(error)
		err1 = errors.New("error 1")

		err2 := errorPool.Get().(error)
		err2 = errors.New("error 2")

		err3 := errorPool.Get().(error)
		err3 = errors.New("error 3")

		err1 = nil
		errorPool.Put(err1)
		err2 = nil
		errorPool.Put(err2)
		err3 = nil
		errorPool.Put(err3)
	}
}

func Benchmark_err_join2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err1 := errorPool.Get().(error)
		err1 = nil
		err1 = errors.New("error 1")

		err2 := errorPool.Get().(error)
		err2 = nil
		err2 = errors.New("error 2")

		err3 := errorPool.Get().(error)
		err3 = nil
		err3 = errors.New("error 3")

		errorPool.Put(err1)
		errorPool.Put(err2)
		errorPool.Put(err3)
	}
}

func Benchmark_err_join3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err1 := errorPool.Get().(*MyError)
		err1.ReSet()
		err1 = &MyError{"err1"}

		err2 := errorPool.Get().(*MyError)
		err2.ReSet()
		err2 = &MyError{"err2"}

		err3 := errorPool.Get().(*MyError)
		err3.ReSet()
		err3 = &MyError{"err3"}

		errorPool.Put(err1)
		errorPool.Put(err2)
		errorPool.Put(err3)
	}
}

func Benchmark_err_join4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err1 := errorPool.Get().(*MyError)
		err1.ReSet()
		err1 = &MyError{"err1"}

		err2 := errorPool.Get().(*MyError)
		err2.ReSet()
		err2 = &MyError{"err2"}

		err3 := errorPool.Get().(*MyError)
		err3.ReSet()
		err3 = &MyError{"err3"}

		errorPool.Put(err1)
		errorPool.Put(err2)
		errorPool.Put(err3)
	}
}

func Benchmark_err_join5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var err MyError
		_ = err
	}
}

func Benchmark_err_join_no_pool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err1 := &MyError{"error 1"}
		err2 := &MyError{"error 2"}
		err3 := &MyError{"error 3"}

		_ = errors.Join(err1, err2, err3)
	}
}
