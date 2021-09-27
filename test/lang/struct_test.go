package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Test struct {
	X int
	Y int
}

func TestStructAssign001(t *testing.T) {
	var g Test

	for i := 0; i < 1000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = Test{1, 2}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = Test{3, 4}
		}()
		wg.Wait()

		// 赋值异常判断
		if !((g.X == 1 && g.Y == 2) || (g.X == 3 && g.Y == 4)) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
			break
		}
	}
}

func TestComplexAssign001(t *testing.T) {
	var g complex64

	for i := 0; i < 1000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = complex(1, 2)
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = complex(3, 4)
		}()
		wg.Wait()

		// 赋值异常判断
		if g != complex(1, 2) && g != complex(3, 4) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
			break
		}
	}
}

func TestStringAssign001(t *testing.T) {
	var s string

	for i := 0; i < 10000000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "ab"
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "abc"
		}()
		wg.Wait()

		// 赋值异常判断
		if s != "ab" && s != "abc" {
			fmt.Printf("concurrent assignment error, i=%v s=%v", i, s)
			break
		}
	}
}

//
// 并发赋值两个等长度但内容不同的字符串，不会有问题
//
func TestStringAssign002(t *testing.T) {
	var s string

	for i := 0; i < 10000000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "123"
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = "abc"
		}()
		wg.Wait()

		// 赋值异常判断
		if s != "123" && s != "abc" {
			fmt.Printf("concurrent assignment error, i=%v s=%v", i, s)
			break
		}
	}
}

type Add func(int, int) int

func TestFuncAssign001(t *testing.T) {
	var g Add

	var i int
	for ; i < 10000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = func(x, y int) int {
				return x + y
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = func(x, y int) int {
				return x - y
			}
		}()
		wg.Wait()

		// 赋值异常判断
		if !(g(1, 1) == 2 || g(1, 1) == 0) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
			break
		}
	}
	if i == 10000000 {
		fmt.Println("no error")
	}
}

func TestArrayAssign001(t *testing.T) {
	var g [4]byte

	var i int
	for ; i < 10000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = [...]byte{1, 2, 3, 4}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = [...]byte{3, 4, 5, 6}
		}()
		wg.Wait()

		// 赋值异常判断
		if !(g == [...]byte{1, 2, 3, 4} || g == [...]byte{3, 4, 5, 6}) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
			break
		}
	}
	if i == 10000000 {
		fmt.Println("no error")
	}
}

func TestArrayAssign002(t *testing.T) {
	var g [5]byte

	var i int
	for ; i < 10000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = [...]byte{1, 2, 3, 4, 5}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = [...]byte{3, 4, 5, 6, 7}
		}()
		wg.Wait()

		// 赋值异常判断
		if !(g == [...]byte{1, 2, 3, 4, 5} || g == [...]byte{3, 4, 5, 6, 7}) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
			break
		}
	}
	if i == 10000000 {
		fmt.Println("no error")
	}
}

type Result struct {
	ID int
	X *BaseResp
	Y *Data
}

type BaseResp struct {
	Code int
}

type Data struct {
	Value int
}

func TestStructAssign002(t *testing.T) {

	batchSize := 13

	for i := 0; i < 100000000000; i++ {
		var wg sync.WaitGroup
		var g []*Result
		for j := 0; j < batchSize; j++ {
			wg.Add(1)

			go func(t int) {

				defer wg.Done()

				time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
				g = append(g, &Result{
					ID: t,
					X: &BaseResp{
						Code: t,
					},
					Y: &Data{
						Value: t,
					},
				})
			}(j)
		}

		wg.Wait()
		// 赋值异常判断

		//if len(g) != batchSize {
		//	bytes, _ := json.Marshal(g)
		//	fmt.Printf("find assignment error, i=%v,len=%v, v=%v", i, len(g), string(bytes))
		//	break
		//} else {
		//	for _, v := range g {
		//		if v != nil && v.X != nil && v.Y != nil {
		//			if v.X.Code != v.Y.Value {
		//				fmt.Printf("find assignment error2222, i=%v v=%+v", i, v)
		//			}
		//		}
		//	}
		//}

		for _, v := range g {
			//if v != nil && v.X != nil && v.Y != nil {
			if v.ID != v.X.Code || v.ID != v.Y.Value || v.X.Code != v.Y.Value {
				fmt.Printf("find assignment error, i=%v v=%+v", i, v)
			}
			//}
		}
	}
}
