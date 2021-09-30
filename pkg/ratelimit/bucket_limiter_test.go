package ratelimit

import (
	"fmt"
	"testing"
	"time"
)

func TestBucketLimiter001(t *testing.T) {
	// 限制每秒2个请求，漏斗容量为5
	r := NewBucketLimiter(2, 5)

	for i := 0; i < 100; i++ {

		ok := r.leakBucket()
		if ok {
			fmt.Println("pass ", i)
		} else {
			fmt.Println("limit ", i)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
