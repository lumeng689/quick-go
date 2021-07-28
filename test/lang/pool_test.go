package test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

/**
 * go test -bench="^Ben"  -cpuprofile=cpu.pprof . 生成profile
 * go tool pprof -http=:9999 ./cpu.pprof
 * go test -bench="^Ben"  -memprofile=mem.pprof .
 * go tool pprof -http=:9999 ./mem.pprof
 */
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func TestSyncPool001(t *testing.T) {

	stu := studentPool.Get().(*Student)
	fmt.Printf("stud01::%v\n", stu)

	json.Unmarshal(buf, stu)
	fmt.Printf("stud02::%v\n", stu)
	studentPool.Put(stu)
	stu3 := studentPool.Get().(*Student)
	fmt.Printf("stud03::%v\n", stu3)

	stu4 := studentPool.Get().(*Student)
	fmt.Printf("stud04::%v\n", stu4)
}

func TestSyncPool002(t *testing.T) {

	studentPool.Put(&Student{
		Name:   "zhangsan",
		Age:    0,
		Remark: [1024]byte{},
	})
	studentPool.Put(&Student{
		Name:   "lisi",
		Age:    0,
		Remark: [1024]byte{},
	})
	studentPool.Put(&Student{
		Name:   "wangwu",
		Age:    0,
		Remark: [1024]byte{},
	})

	stu1 := studentPool.Get().(*Student)
	fmt.Printf("stud01::%v\n", stu1)

	stu2 := studentPool.Get().(*Student)
	fmt.Printf("stud02::%v\n", stu2)

	stu3 := studentPool.Get().(*Student)
	fmt.Printf("stud03::%v\n", stu3)

	stu4 := studentPool.Get().(*Student)
	fmt.Printf("stud04::%v\n", stu4)

	studentPool.Put(stu2)

	stu5 := studentPool.Get().(*Student)
	fmt.Printf("stud05::%v\n", stu5)
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
