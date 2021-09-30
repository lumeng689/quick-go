package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestCounterLimit001(t *testing.T) {
	// 1秒钟最多10个请求
	r := NewCounterLimiter(1, 10)

	for i :=0 ;i < 50;i++ {
		ok := r.counterLimit()
		if ok {
			fmt.Println("pass ", i)
		} else {
			fmt.Println("limit ", i)
		}

		time.Sleep(50*time.Millisecond)
	}
}
