package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println("The time is ", time.Now())
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(32))

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(33))

	var i, j int = 1, 2
	// 另一种声明方式
	k := 3
	fmt.Println(i, j, k, c, python, java)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f1)
	fmt.Println(x, y, z)

	// 常量不能使用 := 语法定义
	const world = "世界"
	fmt.Println("Hello", world)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1
	// C 的 while 在 Go 中叫做 `for`
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// for {}  = while(true){}

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	// switch 的条件从上到下的执行，当匹配成功的时候停止。
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// defer 语句会延迟函数的执行直到上层函数返回。
	// 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
	defer fmt.Println("world")
	fmt.Println("Hello")

	fmt.Println("counting")
	// 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// 以下开始练习Go的指针
	fmt.Println("指针练习")
	i, j = 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1, 2})

	fmt.Println("......End main......")
}

// 连续类型，除最后一个外，其他可以省略
func add(x, y int) int {
	return x + y
}

// 函数可以返回任意数量返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// if else 结构
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}
