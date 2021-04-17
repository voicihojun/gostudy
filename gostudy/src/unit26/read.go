package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// 바이트 배열 20개 크기로 된 buf 슬라이스를 만든다. 
	// file.Read(buf) 를 통해 file, 즉 hello.txt에 담겨진 것을 읽어서 buf 에 담는다.
	buf := make([]byte, 20)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))

}