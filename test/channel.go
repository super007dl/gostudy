package main

import "fmt"

func count(ch chan int) {
	ch <- 1
	fmt.Println("Counting", ch)
}


func main() {
	//创建10个channel数组
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)

		go count(chs[i])
	}

	for _, ch := range(chs) {
		<- ch
	}

}
