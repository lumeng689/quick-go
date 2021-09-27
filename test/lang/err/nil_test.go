package err

import (
	"fmt"
	"testing"
)

func TestNil001(t *testing.T) {

	// 代码中必须提供充足的信息来让编译器推断出某个nil的类型。
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// 下面这一组和上面这一组等价。
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// 下面这行编译不通过。
	var _ = nil

	// error: 类型不匹配
	var _ = (*int)(nil) == (*bool)(nil)
	// error: 类型不匹配
	var _ = (chan int)(nil) == (chan bool)(nil)

	// 这几行编译都没问题。
	var _ = ([]int)(nil) == nil
	var _ = (map[string]int)(nil) == nil
	var _ = (func())(nil) == nil

	fmt.Println((interface{})(nil) == (*int)(nil)) // false

	fmt.Println((map[string]int)(nil)["key"]) // 0
	fmt.Println((map[int]bool)(nil)[123])     // false
	fmt.Println((map[int]*int64)(nil)[123])   //

}

func TestNil002(t *testing.T) {
	for range []int(nil) {
		fmt.Println("Hello")
	}

	for range map[string]string(nil) {
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i)
	}

	for range chan bool(nil) { // 阻塞在此
		fmt.Println("Bye")
	}
}

func TestNil003(t *testing.T) {
	fmt.Println(*new(*int) == nil)         // true
	fmt.Println(*new([]int) == nil)        // true
	fmt.Println(*new(map[int]bool) == nil) // true
	fmt.Println(*new(chan string) == nil)  // true
	fmt.Println(*new(func()) == nil)       // true
	fmt.Println(*new(interface{}) == nil)  // true
}

type Slice []bool

func (s Slice) Length() int {
	return len(s)
}

func (s Slice) Modify(i int, x bool) {
	s[i] = x // panic if s is nil
}

func (p *Slice) DoNothing() {
}

func (p *Slice) Append(x bool) {
	*p = append(*p, x) // 如果p为空指针，则产生一个恐慌。
}

func TestNil004(t *testing.T) {
	// 下面这几行中的选择器不会造成恐慌。
	_ = ((Slice)(nil)).Length
	_ = ((Slice)(nil)).Modify
	_ = ((*Slice)(nil)).DoNothing
	_ = ((*Slice)(nil)).Append

	// 这两行也不会造成恐慌。
	_ = ((Slice)(nil)).Length()
	((*Slice)(nil)).DoNothing()

	// 下面这两行都会造成恐慌。但是恐慌不是因为nil
	// 属主实参造成的。恐慌都来自于这两个方法内部的
	// 对空指针的解引用操作。
	/*
		((Slice)(nil)).Modify(0, true)
		((*Slice)(nil)).Append(true)
	*/
}
