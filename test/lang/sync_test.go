package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSyncCond001(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d) // 模拟一个工作负载
			cond.L.Lock()
			// 下面的修改必须在cond.L被锁定的时候执行
			values[i] = fmt.Sprintf("a%d", i)
			cond.Broadcast() // 可以在cond.L被解锁后发出通知
			cond.L.Unlock()
			// 上面的通知也可以在cond.L未锁定的时候发出。
			//cond.Broadcast() // 上面的调用也可以放在这里
		}(i)
	}

	// 此函数必须在cond.L被锁定的时候调用。
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait() // 必须在cond.L被锁定的时候调用
	}
}

func TestSyncCond002(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	go func() {
		cond.L.Lock()
		go func() {
			cond.L.Lock()
			cond.Broadcast()
			cond.L.Unlock()
		}()
		cond.Wait()
		fmt.Print("a")
		cond.L.Unlock()
		wg.Done()
	}()
	cond.Wait()
	fmt.Print("b")
	cond.L.Unlock()
	wg.Wait()
	fmt.Println("c")
}