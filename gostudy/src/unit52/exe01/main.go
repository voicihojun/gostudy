package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w:= gzip.NewWriter(file)
	// file로 io.Writer 인터페이스를 따르는 인스턴스 w 생성

	defer w.Close()

	s:= "Hello, world!!!!!!"
	w.Write([]byte(s)) // 문자열 압축하여 파일에 저장
	w.Flush()  // 메모리상의 데이터를 파일에 완전히 저장

	// file2, err := os.Open("hello.txt.gz")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file2.Close()

	// r, err := gzip.NewReader(file2)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer r.Close()

	// b, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))
}