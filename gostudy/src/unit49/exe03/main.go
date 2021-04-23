package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"../exe01/hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	fmt.Println(os.O_RDONLY, os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND )
	fmt.Printf("%016b\n", os.O_RDONLY)
	fmt.Printf("%016b\n", os.O_WRONLY)
	fmt.Printf("%016b\n", os.O_RDWR)
	fmt.Printf("%016b\n", os.O_CREATE)
	fmt.Printf("%016b\n", os.O_APPEND)
	fmt.Printf("%016b\n", os.O_APPEND|os.O_WRONLY)
	// 0644 사용자는 읽기 쓰기 / 그룹과 그외는 읽기만 가능

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)
	// 첫번째 매개변수 offset, 두번째 매개변수 기준점
	// os.SEEK_SET : 파일 맨처음을 기준
	// os.SEEK_CUR : 현재 읽기/쓰기 offset 값을 기준
	// os.SEEK_END : 파일 맨끝을 기준

	n, err = file.Read(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}