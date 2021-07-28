package test

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_Caller001(t *testing.T) {
	for i := 0; i < 4; i++ {
		printCaller(i)
	}
}

func printCaller(i int) {
	callSkip(i)
}

func callSkip(i int) {
	pc, file, line, ok := runtime.Caller(i)
	pcName := runtime.FuncForPC(pc).Name()
	fmt.Println(fmt.Sprintf("pc=%v   file=%s   line=%d   recoverable=%t   func=%s", pc, file, line, ok, pcName))
}
