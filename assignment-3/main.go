package main

import (
	"fmt"
	"io"
	"os"
)

type terminalLogger struct{}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	tl := terminalLogger{}

	io.Copy(tl, file)

}

func (terminalLogger) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
