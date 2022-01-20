package mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type node struct {
	ident   string
	context string
}

func run() {
	iter := 100000
	data := make(map[int]node, iter)
	ts := time.Now().String()

	for i := 0; i < iter; i++ {
		data[i] = node{
			ident:   ts,
			context: ts + ts + ts,
		}
	}

	for i := 0; i < iter; i++ {
		delete(data, i)
	}
}

func batchRun() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			run()
			wg.Done()
		}()
	}

	wg.Wait()
}

func Test_MemAlloc001(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		batchRun()
	}

	cost := time.Since(start)
	fmt.Printf("time cost: %+v \n", cost)
}

func Test_MemAlloc002(t *testing.T) {
	start := time.Now()

	cost := time.Since(start)
	fmt.Printf("time cost: %+v \n", cost)
}
