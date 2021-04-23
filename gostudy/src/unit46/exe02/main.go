package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("Hello, world!", "He"))
	fmt.Println(strings.Index("Hello, world!", "wor"))
	fmt.Println(strings.Index("Hello, world!", "ow"))

	fmt.Println(strings.IndexAny("Hello, world!", "eo"))
	fmt.Println(strings.IndexAny("Hello, world!", "f"))

	var c byte
	
	c = 'd'
	fmt.Println(strings.IndexByte("Hello, world!", c))

	c = 'f'
	fmt.Println(strings.IndexByte("Hello, world!", c))

	var r rune = '언'
	
	
	fmt.Println(strings.IndexRune("고 언어", r))
	//한글 1글자는 인덱스 3개를 잡아먹음. 영문 1글자 == 인덱스 1개 
	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	fmt.Println(strings.IndexFunc("Go 언어", f))
	fmt.Println(strings.IndexFunc("Go Language", f))

	fmt.Println(strings.LastIndex("Hello Hello Hello world!", "Hello"))

	fmt.Println(strings.LastIndexAny("Hello, world", "ol"))

}