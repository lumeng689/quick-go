package test

import (
	"fmt"
	"testing"
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
