package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	s:= "Hello, world!"
	h1 := sha512.Sum512([]byte(s))
	// 문자열의 SHA512 해시값 추출
	fmt.Printf("%x\n", h1)	

	sha := sha512.New()
	// SHA512 해시 인스턴스 생성

	sha.Write([]byte("Hello, "))
	sha.Write([]byte("world!"))
	h2 := sha.Sum(nil) // Sum 함수에 nil 넣어주기
	// 해시 인스턴스에 저장된 데이터의 SHA512 해시값 추출
	fmt.Printf("%x\n", h2)


	// 결과값
	// h1, h2 가 같은 값을 나타내야 함. 
	// 만약 sha.Write 부분에 다른 문자열을 넣으면 다른 
	// SHA512 해시값이 리턴됨. 즉 문자열에 따라 해시값 달라짐. 
}