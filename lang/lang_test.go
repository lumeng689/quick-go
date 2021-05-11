package lang

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Alternate_Print001(t *testing.T) {

	flag := make(chan struct{})
	exit := make(chan struct{})

	go func() {
		for i := 1; i < 10; i++ {
			flag <- struct{}{}
			fmt.Println(strconv.Itoa(i))
		}
	}()

	go func() {
		for i := 1; i < 10; i++ {
			<-flag
			fmt.Println(strconv.Itoa(i * 2))
		}

		flag <- struct{}{}
	}()

	<-exit
}
