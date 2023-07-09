package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	data := make([]byte, 4)

	for {
		n, err := file.Read(data)
		if errors.Is(err, io.EOF) {
			break
		}

		fmt.Println(n, string(data[:n]))
	}
}

func readFullHelloFile() {
	fileContent, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(fileContent))
}

func createHelloFile() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello world!\n")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
