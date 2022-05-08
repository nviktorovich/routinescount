package main

import "fmt"

func foo() {
	go zoo()
	go zoo()
	go func() {
		fmt.Println("hello")
	}()
}

func zoo() {
	fmt.Println("hi")
}

func boo() {
	go zoo()
	fmt.Println("No")
}
