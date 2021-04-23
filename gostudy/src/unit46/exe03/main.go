package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := []string{"Hello,", "world!"}
	fmt.Println(strings.Join(s1, " ")) //"Hello, world!"

	s2 := strings.Split("Hello, world!", " ") // [Hello,, world!]
	fmt.Println(s2[1])

	s3 := strings.Fields("Hello, world!") // [Hello,, world!]
	fmt.Println(s3)

	f:= func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	s4 := strings.FieldsFunc("Hello안녕Hello", f)
	fmt.Println(s4) // [Hello Hello]

	fmt.Println(strings.Repeat("Hello", 2)) // "HelloHello"
	fmt.Println(strings.Replace("Hello, world!", "world", "Go", 1)) 
	// Hello, Go!
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))
	// HeGo HeGo
}