package test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_IO001(t *testing.T) {
	sampledata := []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Nunc a mi dapibus, faucibus mauris eu, fermentum ligula.",
		"Donec in mauris ut justo eleifend dapibus.",
		"Donec eu erat sit amet velit auctor tempus id eget mauris.",
	}

	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	dataWriter := bufio.NewWriter(file)

	for _, data := range sampledata {
		_, _ = dataWriter.WriteString(data + "\n")
	}

	dataWriter.Flush()
	file.Close()
}

func TestFile001(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println("current path:", dir)

	stat, err := os.Stat(dir)

	if err != nil {
		if os.IsExist(err) {
			fmt.Println("file exists...")
		} else {
			fmt.Println("file not exits!")
		}
	}

	fmt.Println("is directory: ", stat.IsDir())

	files, err := ioutil.ReadDir(dir)

	for _, v := range files {
		if v.IsDir() {
			fmt.Println("sub directory:", v.Name())
		} else {
			fmt.Println("file:", v.Name())
			fmt.Println("abs path: ", dir+string(os.PathSeparator)+v.Name())
		}
	}
}
