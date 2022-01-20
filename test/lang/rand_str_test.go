package test

import (
	"fmt"
	"testing"
)

func TestRandomString001(t *testing.T) {
	fmt.Println(RandomString(2))
	fmt.Println(RandomString(3))
	fmt.Println(RandomString(6))
}
