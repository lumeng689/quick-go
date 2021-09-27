package test

import (
	"fmt"
	"testing"
)

//
//  go test -v escape_test.go -run Test_Escape001  -gcflags '-m -l'
//
func Test_Escape001(t *testing.T) {
	s := "hello111"
	fmt.Println(s)
}

func Test_Escape002(t *testing.T) {
	b := "ok222"
	fmt.Println(b)
}

type S struct{
	M *int
}

//
//  go test -v escape_test.go -run Test_EscapeExample001  -gcflags '-m -l'
//
func Test_EscapeExample001(t *testing.T) {
	var x S
	y := &x
	_ = *identity(y)
}

func identity(z *S) *S {
	return z
}

func Test_EscapeExample002(t *testing.T) {
	var x S
	_ = *ref(x)
}

func ref(z S) *S {
	return &z
}

func Test_EscapeExample003(t *testing.T) {
	var i int
	refStruct(i)
}

func refStruct(y int) (z S) {
	z.M = &y
	return z
}

func Test_EscapeExample004(t *testing.T) {
	var i int
	refStruct4(&i)
}

func refStruct4(y *int) (z S) {
	z.M = y
	return z
}

func Test_EscapeExample005(t *testing.T) {
	var x S
	var i int
	ref5(&i, &x)
}

func ref5(y *int, z *S) {
	z.M = y
}