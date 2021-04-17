package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello world")
}

func main() {
	fmt.Println("hahahaha 1")
	go hello()

	fmt.Println("hahahaha 2")
	fmt.Scanln()

}