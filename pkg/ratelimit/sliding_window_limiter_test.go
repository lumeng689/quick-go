package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestSlidingWindow001(t *testing.T) {
	// 定义1秒10个时间窗口，每个窗口大小为1， 即 1秒10个请求
	r := NewSlidingWindowLimiter(1, 10, 1)
	for i := 0; i < 300; i++ {
		ok := r.slidingCounterLimit()
		if ok {
			fmt.Println("pass ", i)
		} else {
			fmt.Println("limit ", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
