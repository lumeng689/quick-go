package main

import (
	"fmt"
	"runtime"
)

func example001(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}

//
// 查看编译的时候内连情况 go build -gcflags "-m" panic1.go
// 禁止内连的编译go代码  go build -gcflags "-l" panic1.go
func main() {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Println(string(buf))
		}
	}()
	example001(true, false, true, 25)
}
