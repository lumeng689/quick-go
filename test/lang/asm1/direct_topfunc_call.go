package main

//go:noinline
func add(a, b int32) (int32, bool) { return a + b, true }

// 执行指令， 将代码编译为汇编形式
// GOOS=linux GOARCH=amd64 go tool compile -S direct_topfunc_call.go
func main() { add(10, 32) }
