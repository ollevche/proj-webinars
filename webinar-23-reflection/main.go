package main

import (
	"fmt"
	"time"
)

func main() {
	entity := &DTONewEan{}

	tags, err := GetTagsJSON(entity)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tags)

	var g jsonTagGetter

	tags, err = g.Get(entity)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tags)

	tags, err = g.Get(entity)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tags)
}

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recovered: %v\n", r)
		}
	}()

	startedAt := time.Now()
	defer func() {
		fmt.Printf("Execution took %v\n", time.Since(startedAt))
	}()

	// panic("Something went wrong")

	startNewGoroutine()

	time.Sleep(3 * time.Second)

	fmt.Println("Hello")
}

func startNewGoroutine() {
	go func() {
		defer func() {
			fmt.Println("Hello from startNewGoroutine")
		}()

		childFunc()
	}()
}

func childFunc() {
	defer func() {
		fmt.Println("Hello from childFunc")
	}()

	time.Sleep(2 * time.Second)
	panic("Can't do it")
}
