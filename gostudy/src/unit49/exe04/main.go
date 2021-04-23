package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := "Hello, world!!!!!"
	err := ioutil.WriteFile("../exe01/hello.txt", []byte(s), os.FileMode(644))
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile("../exe01/hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}