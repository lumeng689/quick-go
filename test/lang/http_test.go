package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type Message struct {
	Code    int32       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Test_HttpGet001(t *testing.T) {

	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	var query = url.Values{}
	query.Add("q", "你好")

	req.URL.RawQuery = query.Encode()

	c := http.DefaultClient
	rsp, err := c.Do(req)

	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		panic(err)
	}

	if msg.Data == nil {
		fmt.Printf("no data... %v \n", string(body))
	} else {
		fmt.Printf("data is:%v 、n", msg.Data)
	}
}

func Test_HttpGet002(t *testing.T) {

	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	c := http.DefaultClient

	for i := 0; i < 500; i++ {
		rsp, err := c.Do(req)

		if err != nil {
			panic(err)
		}
		defer rsp.Body.Close()
	}
}
