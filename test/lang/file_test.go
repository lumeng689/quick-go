package test

import (
	"testing"
)

func TestWriteStringToFile001(t *testing.T) {
	WriteStringToFile("/tmp/aa.txt", "hello world")
}
