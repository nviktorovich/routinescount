package main

import (
	"fmt"
)

func recount() {
	return
}

func count(i, z int) {
	fmt.Println(i, z)
	recount()
	go recount()
	recount()
	go recount()
	go recount()
	go func() {
		fmt.Println("ooo")
	}()
}

func Stop() {
	fmt.Println("no no no")
}
