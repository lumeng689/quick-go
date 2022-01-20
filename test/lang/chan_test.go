package test

import (
	"fmt"
	"testing"
)

//
// for 循环，读取10次之后，自动退出了
//
func TestReadChan001(t *testing.T) {
	count := 10
	ch := generate(count)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func generate(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			ch <- i
		}
	}()
	return ch
}

//
// for range 循环，可能会导致挂起，只有在channel关闭之后， range循环才结束。
//
func TestReadChan002(t *testing.T) {
	count := 10
	ch := generate(count)
	for v := range ch {
		fmt.Println(v)
	}
}

func TestCloseChan001(t *testing.T) {
	c := make(chan int, 2) // 一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)

	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2

	close(c) // 此行将产生一个恐慌
	c <- 7   // 如果上一行不存在，此行也将产生一个恐慌。
}
