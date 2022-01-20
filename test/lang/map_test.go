package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestMap001(t *testing.T) {
	dataMap := make(map[int]int)

	for i := 0; i < 1000; i++ {
		dataMap[i] = 0
	}

	for i := 0; i < 1000; i++ {
		go func(i int) {
			dataMap[i] = i * 100
		}(i)
	}

	time.Sleep(1000 * time.Second)
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
	runtime.GC() // 手动触发 GC
	fmt.Printf("With a map of strings, GC took: %s\n", time.Since(now))

	_ = m["0"] // 引用一下防止被 GC 回收掉
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

func TestMapNestSlice001(t *testing.T) {
	aaa := map[int][]int64{}

	b := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3}

	for _, v := range b {
		list, ok := aaa[v]
		if !ok {
			list = make([]int64, 0, 5)
			aaa[v] = list
		}
		list = append(list, int64(v))
		aaa[v] = list
	}
	fmt.Printf("aaa: %v \n", aaa)

	for _, v := range aaa {
		s := pickRandom(v, 2)
		fmt.Printf("s:%v \n", s)
	}

}

func pickRandom(list []int64, count int) []int64 {
	if count <= 0 || len(list) <= count {
		return list
	}

	ret := make([]int64, 0, count)
	size := len(list)
	startIdx := rand.Intn(size)
	for i := 0; i < count; i++ {
		if startIdx >= size {
			startIdx = startIdx % size
		}
		ret = append(ret, list[startIdx])
		startIdx++
	}

	return ret
}

func TestMapNestSlice002(t *testing.T) {
	aaa := make(map[int]int64)

	b := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3}

	for _, v := range b {
		aaa[v]++
	}

	fmt.Printf("%v \n", aaa)
}

func TestMapNestSlice003(t *testing.T) {
	for i := 0; i < 100; i++ {
		startIdx := rand.Intn(3)
		fmt.Printf("%v\n", startIdx)
	}

}

func TestMapNestSlice004(t *testing.T) {
	aaa := make(map[int]int64)
	aaa[1] = 1
	aaa[2] = 2
	fmt.Printf("map:%v \n", aaa)
	delete(aaa, 2)
	delete(aaa, 3)
	fmt.Printf("map:%v \n", aaa)
}

func TestMapRange001(t *testing.T) {
	countMap := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		countMap[i] = i * 2
	}

	list := make([]*int, 0)

	for _, v := range countMap {
		list = append(list, &v)
	}

	for _, v := range list {
		fmt.Println(v)
	}
}
