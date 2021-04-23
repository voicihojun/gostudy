package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))
	fmt.Println(strings.Contains("Hello, world!", "w o"))
	fmt.Println(strings.Contains("Hello, world!", "ow"))

	fmt.Println("===========")

	fmt.Println(strings.ContainsAny("Hello,world!", "wo"))
	fmt.Println(strings.ContainsAny("Hello,world!", "w o")) 
	// 검색해야 할 문자열에 공백이 없어도 "w o"에 대해서 true 값 반환함
	fmt.Println(strings.ContainsAny("Hello,world!", "ow"))

	fmt.Println("===========")

	fmt.Println(strings.Count("Hello, Helium!", "He"))

	var r rune
	r = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r))
	fmt.Println(strings.HasPrefix("Hello, world!", "He"))
	fmt.Println(strings.HasSuffix("Hello, world!", "rld!"))


}