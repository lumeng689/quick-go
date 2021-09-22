package test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func TestSlice001(t *testing.T) {

	arr := make([]int, 5)

	for i := 3; i < 9; i++ {
		arr = append(arr, i)
	}

	for i, v := range arr {
		fmt.Printf("%d, %d \n", i, v)
	}
}

func TestSlice002(t *testing.T) {

	arr := make([]int, 0)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			mu.Lock()
			defer mu.Unlock()
			arr = append(arr, j)
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i, v := range arr {
		fmt.Printf("%d, %d \n", i, v)
	}
}

func TestSlice003(t *testing.T) {

	arr := make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		arr = append(arr, i)
	}

	for i, v := range arr {
		fmt.Printf("%d, %d \n", i, v)
	}
}

func TestSlice004(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("%d\n", i)
		}()
	}

	wg.Wait()
}


func TestStress(t *testing.T) {
	iters := 1000000
	if testing.Short() {
		iters = 1000
	}
	var sink []byte
	for i := 0; i < iters; i++ {
		//_ = Get("foo")
		sink = make([]byte, 1<<20)
	}
	_ = sink
}

func aaa() {
	var s []string


	for i := 0;i < 64;i++ {
		go func() {
			s = append(s, "a")
		}()
	}

	time.Sleep(5 * time.Second)
}

func TestStress001(t *testing.T) {
	aaa()
}

type MyData struct {
	aByte   byte
	//anotherByte byte
	aShort  int16
	anInt32 int32
	aSlice  []byte
}

func TestMemLayout001(t *testing.T) {
	typ := reflect.TypeOf(MyData{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())

	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
}

func TestMemLayout002(t *testing.T) {
	data := MyData{
		aByte:   0x1,
		aShort:  0x0203,
		anInt32: 0x04050607,
		aSlice:  []byte{
			0x08, 0x09, 0x0a,
		},
	}

	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	fmt.Printf("Bytes are %#v\n", dataBytes)

	dataslice := *(*reflect.SliceHeader)(unsafe.Pointer(&data.aSlice))
	fmt.Printf("Slice data is %#v\n",
		(*[3]byte)(unsafe.Pointer(dataslice.Data)))
}