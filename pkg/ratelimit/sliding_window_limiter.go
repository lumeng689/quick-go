package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	Interval       int64       // 总计数时间
	LastTimeSec    int64       // 上一个窗口时间
	Lock           *sync.Mutex // 锁
	WinCount       []int64     // 窗口计数器，记录每个窗口的请求数量
	TicketSize     int64       // 窗口最大容量
	TicketCount    int64       // 窗口个数
	EachTicketTime int64       // 每个窗口的时间
	CurIndex       int64       // 目前使用的窗口
}

func NewSlidingWindowLimiter(interval int64, ticketCount int64, ticketSize int64) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		Interval:       interval,
		LastTimeSec:    time.Now().Unix(),
		TicketSize:     ticketSize,
		TicketCount:    ticketCount,
		EachTicketTime: interval / ticketCount,
		WinCount:       make([]int64, ticketCount),
		CurIndex:       0,
		Lock:           new(sync.Mutex),
	}
}

func (r *SlidingWindowLimiter) slidingCounterLimit() bool {
	r.Lock.Lock()
	defer r.Lock.Unlock()

	nowSec := time.Now().Unix()

	if nowSec-r.LastTimeSec > r.EachTicketTime {
		r.WinCount[r.CurIndex] = 0
		r.CurIndex = (r.CurIndex + 1) % r.TicketCount
		r.LastTimeSec = nowSec
	}

	fmt.Println("当前窗口: ", r.CurIndex)

	if r.WinCount[r.CurIndex] < r.TicketSize {
		r.WinCount[r.CurIndex]++
		return true
	}

	return false
}
