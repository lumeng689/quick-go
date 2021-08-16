package test

import (
	"fmt"
	"reflect"
	"testing"
)

func myFunc(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	switch t.Kind() {
	case reflect.Int:
		fmt.Println(1 + v.Int())
	case reflect.String:
		fmt.Println("hello " + v.String())
	default:
		panic("x is not a supported type!")
	}
}

func Test_MyFunc001(t *testing.T) {
	myFunc(1)
	myFunc("1")

	a := 3
	myFunc(&a)
}
