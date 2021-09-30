package ratelimit

import (
	"sync"
	"time"
)

type BucketLimiter struct {
	Rate     float64
	Balance  float64
	Cap      float64
	LastTime time.Time
	Lock     *sync.Mutex
}

func NewBucketLimiter(rate, cap int) *BucketLimiter {
	return &BucketLimiter{
		Rate:     float64(rate),
		Balance:  float64(cap),
		Cap:      float64(cap),
		LastTime: time.Now(),
		Lock:     new(sync.Mutex),
	}
}

func (r *BucketLimiter) leakBucket() bool {
	ok := false

	r.Lock.Lock()
	defer r.Lock.Unlock()

	now := time.Now()
	// sub一下，然后seconds，出来的结果是float64的，方便后面参加运算
	diff := now.Sub(r.LastTime).Seconds()
	r.LastTime = now

	water := diff * r.Rate
	r.Balance += water
	if r.Balance > r.Cap {
		r.Balance = r.Cap
	}

	if r.Balance >= 1 {
		r.Balance--
		ok = true
	}

	return ok
}
