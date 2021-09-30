package ratelimit

import (
	"fmt"
	"github.com/juju/ratelimit"
	"sync"
	"testing"
	"time"
)

func TestTokenBucket000(t *testing.T) {
	tb := ratelimit.NewBucket(50*time.Millisecond, 10)

	for i := 0; i < 20; i++ {
		// 取token 非阻塞，返回值是等待时间
		d := tb.Take(1)
		fmt.Printf("%v: %v, duration: %v\n", time.Now().Unix(), i, d)
	}

	tb.TakeMaxDuration(1, 5*time.Second)
}

func TestTokenBucket001(t *testing.T) {
	tb := ratelimit.NewBucket(50*time.Millisecond, 10)

	for i := 0; i < 100; i++ {
		// 取token 非阻塞，返回值是拿到的token数
		a := tb.TakeAvailable(3)
		fmt.Printf("%v: %v, available: %v\n", time.Now().Unix(), i, a)
	}

	tb.TakeMaxDuration(1, 5*time.Second)
}

func TestTokenBucket002(t *testing.T) {
	tb := ratelimit.NewBucket(1*time.Second, 10)

	for i := 0; i < 100; i++ {
		// 最多等maxWait时间取token
		d, ok := tb.TakeMaxDuration(1, 500*time.Second)
		if ok {
			fmt.Printf("%v: %v, duration: %v\n", time.Now().Unix(), i, d)
		} else {
			fmt.Println("wait too long....")
		}
	}
}

func TestTokenBucket003(t *testing.T) {
	tb := ratelimit.NewBucket(500*time.Millisecond, 10)

	for i := 0; i < 20; i++ {
		// 取token（阻塞）
		tb.Wait(1)
		fmt.Printf("%v: %v\n", time.Now().Unix(), i)
	}
}

func TestTokenBucket004(t *testing.T) {
	tb := ratelimit.NewBucket(100*time.Millisecond, 10)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			ok := tb.WaitMaxDuration(1, 500*time.Second)
			if ok {
				fmt.Printf("%v: %v\n", time.Now().Unix(), i)
			} else {
				fmt.Println("wait too long")
			}
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}

func TestTokenBucket005(t *testing.T) {
	// 创建指定填充速率和容量大小的令牌桶
	_ = ratelimit.NewBucket(1*time.Second, 10)
	// 创建指定填充速率、容量大小、 每次填充的令牌数的令牌桶
	_ = ratelimit.NewBucketWithQuantum(1*time.Second, 50, 2)
	// NewBucketWithRate(0.1, 200) 表示每秒填充20个令牌
	_ = ratelimit.NewBucketWithRate(0.1, 200)
}
