package test

import (
	"fmt"
	"net/url"
	"testing"
)

func TestEscape001(t *testing.T) {

	var s = `device_brand=iPhone X&label=open_sv_draw_cheers&gy=d11d58`

	fmt.Println(url.QueryEscape(s))
}
