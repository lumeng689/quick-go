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

func TestUrlBuilder001(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	u.Opaque = "opaque"
	fmt.Println(u.String())
}

func TestUrlBuilder002(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}

	fmt.Println(u.String())
}
