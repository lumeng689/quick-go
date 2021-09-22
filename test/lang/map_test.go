package test

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestMap001(t *testing.T) {
	dataMap := make(map[int]int)

	for i := 0; i < 5; i++ {
		dataMap[i] = i * 100
	}

	for _, v := range dataMap {
		fmt.Printf("====%v\n", v)
	}
}

func MapWithPointer() {
	const N = 10000000
	m := make(map[string]string)
	for i := 0; i < N; i++ {
		n := strconv.Itoa(i)
		m[n] = n
	}
	now := time.Now()
	runtime.GC()      // 手动触发 GC
	fmt.Printf("With a map of strings, GC took: %s\n", time.Since(now))

	_ = m["0"]          // 引用一下防止被 GC 回收掉
}

func MapWithoutPointer() {
	const N = 10000000
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		str := strconv.Itoa(i)
		// hash string to int
		n, _ := strconv.Atoi(str)
		m[n] = n
	}
	now := time.Now()
	runtime.GC()
	fmt.Printf("With a map of int, GC took: %s\n", time.Since(now))

	_ = m[0]
}

func TestMapWithPointer(t *testing.T) {
	MapWithPointer()
}

func TestMapWithoutPointer(t *testing.T) {
	MapWithoutPointer()
}