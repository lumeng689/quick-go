package ratelimit

import (
	"sync"
	"time"
)

type CounterLimiter struct {
	Interval    int64       // 计时周期 单位秒
	LastTimeSec int64       // 上一次计时时间
	MaxCount    int         // 允许的最大数量
	Lock        *sync.Mutex // 锁，保证操作
	ReqCount    int         // 当前计时周期总访问数
}

func NewCounterLimiter(interval int64, maxCount int) *CounterLimiter {
	return &CounterLimiter{
		Interval:    interval,
		LastTimeSec: time.Now().Unix(),
		MaxCount:    maxCount,
		Lock:        new(sync.Mutex),
		ReqCount:    0,
	}
}

func (r *CounterLimiter) counterLimit() bool {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	nowSec := time.Now().Unix()

	if nowSec-r.LastTimeSec > r.Interval {
		r.LastTimeSec = nowSec
		r.ReqCount = 0
	}

	if r.ReqCount < r.MaxCount {
		r.ReqCount += 1
		return true
	}

	return false
}
