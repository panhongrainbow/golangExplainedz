package example

import "sync"

// 接口 I
type I interface {
	Set(x int)
}

type T struct {
	x int
}

func (t *T) Set(x int) {
	t.x = x // 修改接口变量指向的值
}

var obj I

func write() {
	t := new(T)
	obj = t // 修改接口变量
}

func read() {
	obj.Set(10) // 访问接口变量指向的值
}

var mu sync.Mutex

func writeMu() {
	t := new(T)
	mu.Lock()
	obj = t // 修改接口变量
	mu.Unlock()
}

func readMu() {
	mu.Lock()
	obj.Set(10) // 访问接口变量指向的值
	mu.Unlock()
}
