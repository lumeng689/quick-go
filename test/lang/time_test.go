package test

import (
	"fmt"
	"testing"
	"time"
)

func TestDate001(t *testing.T) {
	//var cstZone = time.FixedZone("CST", 8*3600) // 东八
	d := time.Date(1949, time.October, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("unix seconds:", d.Unix())

	fmt.Println("now unix:", time.Now().Unix())
}

func TestDate002(t *testing.T) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
}

func TestDate003(t *testing.T) {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	fmt.Println("SH : ", time.Now().In(cstZone).Format("2006-01-02 15:04:05"))
}

func TestDuration001(t *testing.T) {

	d1 := 100 * time.Second
	d2 := 1 * time.Hour

	fmt.Println(d1 < d2)

}

type FuncA func()

func TestAAA(t *testing.T) {

	tt := []int{1, 2, 3, 4, 5, 6, 7}

	ff := make([]FuncA, 0, len(tt))

	for i, _ := range tt {
		f := func() {
			fmt.Println(i)
		}

		ff = append(ff, f)
	}

	for _, v := range ff {
		v()
	}
}
