package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func showMsg(i int) {
	//wg.Done == wg.Add(-1)
	defer wg.Done()
	//defer 延迟执行 最后执行
	fmt.Println("第  ", i, "  次")
}
func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("主进程.end")
}

