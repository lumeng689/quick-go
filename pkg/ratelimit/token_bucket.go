package ratelimit

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity float64
	Rate     float64
	Tokens   float64
	LastTime time.Time
	Lock     *sync.Mutex
}

func NewTokenBucket(rate, cap int) *TokenBucket {
	return &TokenBucket{
		Capacity: float64(cap),
		Rate:     float64(rate),
		Tokens:   float64(cap),
		Lock:     new(sync.Mutex),
	}
}

func (r *TokenBucket) getToken() bool {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	now := time.Now()

	tokens := math.Min(r.Capacity, r.Tokens+now.Sub(r.LastTime).Seconds()*r.Rate)
	r.Tokens = tokens

	if tokens < 1 {
		return false
	}

	r.Tokens--
	r.LastTime = now
	return true
}
