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
