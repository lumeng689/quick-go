package test

import (
	"fmt"
	"testing"
	"unsafe"
)

type notifier interface {
	notify()
}

func Test_Inf001(t *testing.T) {
	var inf notifier
	// 接口变量是一个占有两个字长度的数据结构，
	// 第一个字包含一个指向内部表的指针，叫iTable，iTable包含了所存储变量的类型信息以及和这个变量相关的方法集。
	// 第二个字是一个指向所存储变量的指针。
	fmt.Println("interface bytes: ", unsafe.Sizeof(inf))
}
