package third

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func Test_JsonMarshal001(t *testing.T) {
	data := map[string]string{"a": "b"}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(&data)

	if err != nil {
		fmt.Printf("error is: %v \n", err)
	}

	fmt.Printf("json string: %v \n", string(bytes))
}

func Test_JsonUnmarshal001(t *testing.T) {
	var data map[string]string
	var input = `{"a": "b"}`
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal([]byte(input), &data)

	fmt.Printf("%v \n", data)
}
