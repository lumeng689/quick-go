package test

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*")

// RandomString 获取随机字符串
func RandomString(n int) string {
	r := make([]rune, n)

	for i := 0; i < n; i++ {
		r[i] = letters[rand.Intn(len(letters))]
	}

	return string(r)
}
