package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestTokenBucket001(t *testing.T) {
	// 初始化 限制每秒2个请求 令牌桶容量为5
	r := NewTokenBucket(2, 5)
	for i := 0; i < 20; i++ {
		ok := r.getToken()
		if ok {
			fmt.Println("pass ", i)
		} else {
			fmt.Println("limit ", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
