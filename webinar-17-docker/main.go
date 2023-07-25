package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	wordToPrint := os.Getenv("WORD_TO_PRINT")

	fmt.Println(wordToPrint)

	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)

		fmt.Println("Working hard... + " + wordToPrint)
	}

	fmt.Println("Done!", wordToPrint)
}
