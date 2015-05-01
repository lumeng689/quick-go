package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
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

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	var v1 = Vertex{X: 1}
	fmt.Println(v1.X)

	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "world"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	pa := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("pa[%d] == %d\n", i, pa[i])
	}

	fmt.Println("pa[1:4] ==", pa[1:4])

	arr1 := make([]int, 5)
	printSlice("arr1", arr1)
	arr2 := make([]int, 0, 5)
	printSlice("arr2", arr2)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	var m map[string]Vertex
	m = make(map[string]Vertex)

	m["Bell labs"] = Vertex{1, 3}
	fmt.Println(m["Bell labs"])

	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println("m2...", m2["Google"])

	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	// 通过双赋值检测某个键存在
	// 如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	v11, ok1 := m3["Answer"]
	fmt.Println("The value:", v11, "Present?", ok1)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	v4 := &Vertex{3, 4}
	fmt.Println(v4.Abs())

	if err := run2(); err != nil {
		fmt.Println(err)
	}

	iii, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	}
	fmt.Println("Converted integer:", iii)

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

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func run2() error {
	return &MyError{time.Now(), "it didn't work"}
}
