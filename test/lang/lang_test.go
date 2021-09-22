package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_aa(t *testing.T) {
	ids := make([]int64, 10)

	for i := 0; i < 10; i++ {
		ids[i] = int64(i)
	}

	for _, v := range ids {
		fmt.Println(v)
	}
}

type Person struct {
	Name   string
	Age    int
	Gender bool
}

// Sizeof返回类型v本身数据所占用的字节数
func Test_SizeOf(t *testing.T) {
	var i int
	fmt.Println("int bytes: ", unsafe.Sizeof(i))

	var i32 int32
	fmt.Println("int32 bytes: ", unsafe.Sizeof(i32))

	var i64 int64
	fmt.Println("int64 bytes: ", unsafe.Sizeof(i64))

	var f32 float32
	fmt.Println("float32 bytes: ", unsafe.Sizeof(f32))

	var f64 float64
	fmt.Println("float64 bytes: ", unsafe.Sizeof(f64))

	var s string
	fmt.Println("string bytes: ", unsafe.Sizeof(s))

	var a []int
	// 返回的是切片描述符的大小
	fmt.Println("[]int a bytes: ", unsafe.Sizeof(a))

	a2 := make([]int, 0)
	a2 = append(a2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("len a2: ", len(a2))
	fmt.Println("cap a2: ", cap(a2))
	fmt.Println("[]int a2 bytes: ", unsafe.Sizeof(a2))

	p1 := new(Person)
	fmt.Println("person1: ", unsafe.Sizeof(p1))

	p2 := Person{}
	fmt.Println("person2: ", unsafe.Sizeof(p2))

	m1 := make(map[int]int, 0)
	fmt.Println("map1: ", unsafe.Sizeof(m1))

	m2 := make(map[string]int, 0)
	fmt.Println("map2: ", unsafe.Sizeof(m2))
}

func Test_SizeOf02(t *testing.T) {
	var intValue int
	var uint8Value uint8
	var uint16Value uint16
	var uint32Value uint32
	var uint64Value uint64
	var int8Value int8
	var int16Value int16
	var int32Value int32
	var int64Value int64

	var float32Value float32
	var float64Value float64
	var boolValue bool
	var ptrValue uintptr
	var complex64Value complex64
	var complex128Value complex128
	var strValue string
	var byteValue byte
	var runeValue rune
	structValue := struct {
		FieldA float32
		FieldB string
	}{0, ""}
	mapValue := map[int]int{}
	var sliceValue []int
	var intPtrValue *int
	var chanValue chan int
	var funcValue func()

	fmt.Println("funcValue = Size:", unsafe.Sizeof(funcValue))     //size:  8
	fmt.Println("chanValue = Size:", unsafe.Sizeof(chanValue))     //size:  8
	fmt.Println("intPtrValue = Size:", unsafe.Sizeof(intPtrValue)) //size:  8
	fmt.Println("sliceValue = Size:", unsafe.Sizeof(sliceValue))   //size:  24
	//type slice struct {
	//    array unsafe.Pointer
	//    len   int
	//    cap   int
	//}
	fmt.Println("mapValue = Size:", unsafe.Sizeof(mapValue))                //size:  8
	fmt.Println("structValue = Size:", unsafe.Sizeof(structValue))          //size:  24
	fmt.Println("strValue = Size:", unsafe.Sizeof(strValue), len(strValue)) //intValue = Size: 16, string is the set of all strings of 8-bit bytes, conventionally but not
	// necessarily representing UTF-8-encoded text. A string may be empty, but
	// not nil. Values of string type are immutable.
	//type stringStruct struct {
	//    str unsafe.Pointer
	//    len int
	//}
	fmt.Println("byteValue = Size:", unsafe.Sizeof(byteValue))             //intValue = Size: 1
	fmt.Println("runeValue = Size:", unsafe.Sizeof(runeValue))             //intValue = Size: 4
	fmt.Println("boolValue = Size:", unsafe.Sizeof(boolValue))             //intValue = Size: 1
	fmt.Println("ptrValue = Size:", unsafe.Sizeof(ptrValue))               //intValue = Size: 8
	fmt.Println("complex64Value = Size:", unsafe.Sizeof(complex64Value))   //intValue = Size: 8
	fmt.Println("complex128Value = Size:", unsafe.Sizeof(complex128Value)) //intValue = Size: 16
	fmt.Println("intValue = Size:", unsafe.Sizeof(intValue))               //intValue = Size: 8
	fmt.Println("uint8Value = Size:", unsafe.Sizeof(uint8Value))           //uint8Value = Size: 1
	fmt.Println("uint16Value = Size:", unsafe.Sizeof(uint16Value))         //uint16Value = Size: 2
	fmt.Println("uint32Value = Size:", unsafe.Sizeof(uint32Value))         //uint32Value = Size: 4
	fmt.Println("uint64Value = Size:", unsafe.Sizeof(uint64Value))         // uint64Value = Size: 8

	fmt.Println("int8Value = Size:", unsafe.Sizeof(int8Value))   //int8Value = Size: 1
	fmt.Println("int16Value = Size:", unsafe.Sizeof(int16Value)) //int16Value = Size: 2
	fmt.Println("int32Value = Size:", unsafe.Sizeof(int32Value)) //int32Value = Size: 4
	fmt.Println("int64Value = Size:", unsafe.Sizeof(int64Value)) //int64Value = Size: 8

	fmt.Println("float32Value = Size:", unsafe.Sizeof(float32Value)) //float32Value = Size: 4
	fmt.Println("float64Value = Size:", unsafe.Sizeof(float64Value)) //float64Value = Size: 8

}

func Test_Pointer001(t *testing.T) {
	var m map[string]int
	m = map[string]int{"one": 1, "two": 2}
	n := m
	fmt.Printf("%p\n", &m) //0xc000074018
	fmt.Printf("%p\n", &n) //0xc000074020
	fmt.Println(m)         // map[two:2 one:1]
	fmt.Println(n)         //map[one:1 two:2]
	changeMap001(m)
	fmt.Printf("%p\n", &m) //0xc000074018
	fmt.Printf("%p\n", &n) //0xc000074020
	fmt.Println(m)         //map[one:1 two:2 three:3]
	fmt.Println(n)         //map[one:1 two:2 three:3]
}
func changeMap001(m map[string]int) {
	m["three"] = 3
	fmt.Printf("changeMap func %p\n", m) //changeMap func 0xc000060240
}

func Test_Pointer002(t *testing.T) {
	var m map[string]int
	m = map[string]int{"one": 1, "two": 2}
	n := m
	fmt.Printf("%p\n", &m) //0xc000074018
	fmt.Printf("%p\n", &n) //0xc000074020
	fmt.Println(m)         // map[two:2 one:1]
	fmt.Println(n)         //map[one:1 two:2]
	changeMap002(&m)
	fmt.Printf("%p\n", &m) //0xc000074018
	fmt.Printf("%p\n", &n) //0xc000074020
	fmt.Println(m)         //map[one:1 two:2 three:3]
	fmt.Println(n)         //map[two:2 three:3 one:1]
}

func changeMap002(m *map[string]int) {
	//m["three"] = 3 //这种方式会报错 invalid operation: m["three"] (type *map[string]int does not support indexing)
	(*m)["three"] = 3                    //正确
	fmt.Printf("changeMap func %p\n", m) //changeMap func 0x0
}
