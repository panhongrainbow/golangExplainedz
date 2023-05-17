package race

import (
	"os"
	"sync"
	"testing"
)

func writeFile(path string, content string) {
	file, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	file.WriteString(content)
}

func Test_Race_iocopy(t *testing.T) {
	file, _ := os.Create("file.txt")
	defer file.Close()

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {

		go func() {
			wg.Done()
			file.WriteString("a")
		}()
	}

	wg.Wait()
}

/*
go


func main() {
    file, _ := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE, 0644)
    defer file.Close()

    go writeFile("file.txt", "hello")
    go writeFile("file.txt", "world")

    time.Sleep(2 * time.Second)  // 等待2秒,让goroutine执行完

    data, _ := ioutil.ReadFile("file.txt")
    fmt.Println(string(data))  // 打印最终文件内容
}
*/
