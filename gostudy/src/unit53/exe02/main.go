package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key:= "Hello, Key 12345" // 16바이트
	s := "Hello, world! 12" // 16바이트
	// AES는 블록 암호화 알고리즘으로 키와 암호화할 데이터의 
	// 크기가 일정해야 함. 여기서는 16바이트 
	
	block, err := aes.NewCipher([]byte(key))
	// AES 대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return 
	}

	ciphertext := make([]byte, len(s))
	// ciphertext에 슬라이스 생성
	block.Encrypt(ciphertext, []byte(s))
	// ciphertext 슬라이스에 s 문자열을 암호화해서 집어넣음
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext)
	// plaintext 슬라이스에 암호화된 ciphertext를 집어넣어서 해석함

	fmt.Println(string(plaintext))
}