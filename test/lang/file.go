package test

import (
	"bufio"
	"fmt"
	"os"
)

func WriteStringToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("failed creating file:%v, error: %v \n", filename, err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Printf("failed writing file:%v, error: %v \n", filename, err)
		return
	}

	writer.Flush()
}
