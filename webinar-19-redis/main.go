package main

import (
	"fmt"

	"github.com/mediocregopher/radix/v3"
)

// username:messageID message_text

func main() {
	p, err := radix.NewPool("tcp", "127.0.0.1:6379", 3)
	if err != nil {
		panic(err.Error())
	}

	err = p.Do(radix.Cmd(nil, "SETNX", "user:1:name", "doNotOverride"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done writing!")

	var userName string

	err = p.Do(radix.Cmd(&userName, "GET", "user:1:name"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done reading!", userName)

	err = p.Do(radix.Cmd(nil, "DEL", "user:1:name"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done deleting!")

	err = p.Do(radix.Cmd(nil, "SETEX", "user:1:counter", "30", "1"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done setting counter!")

	var counter int

	err = p.Do(radix.Cmd(&counter, "INCR", "user:1:counter"))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Done getting counter!", counter)
}
