package main

import (
	"sync"
	"fmt"
	"runtime"
)

//初始化计数器
var counter int = 0

func count(lock *sync.Mutex) {
	//加锁
	lock.Lock()

	counter++

	fmt.Println(counter)
	//释放锁
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i< 10; i++ {
		go count(lock)
	}

	for {
		lock.Lock()

		c := counter

		lock.Unlock()

		runtime.Gosched()

		if c >= 10 {
			break
		}
	}


}
