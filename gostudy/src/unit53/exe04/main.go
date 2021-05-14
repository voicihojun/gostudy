package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	// 2048은 키의 길이 : (키의 길이는 비트단위이며 2의 제곱수로 입력)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey
	//위에서 만든 privateKey 변수에 공개키도 들어있음

	s:= `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	ciphertext, _ := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s),
	)

	fmt.Printf("%x\n", ciphertext)

	plaintext, _ := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext,
	)

	fmt.Println(string(plaintext))
}