package test

import (
	"bufio"
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
