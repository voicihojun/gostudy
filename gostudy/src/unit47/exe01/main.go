package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error
	var num1 int
	num1, err = strconv.Atoi("100")

	fmt.Println(num1, err)

	var num2 int
	num2, err = strconv.Atoi("10t")

	fmt.Println(num2, err)

	s := strconv.Itoa(50)
	fmt.Println(s)
}