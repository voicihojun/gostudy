package main

import (
	"fmt"

	"hojun.com/m/src/unit39/calc"
)

// src 이하 부분 경로를 import 하면 된다고 책에 나와 있는데,
// 내 GOPATH 가 잘못된 것인지 unit39/calc is not in GOROOT 라고 나옴
// /Users/hojunhwang/Documents/gitfth/golang/gostudy/src/unit39
// 위와 같이 전체 경로를 적고, 저장하니 현재 위에 경로가 나옴.
// 처음에 설정했던 go mod init hojun.com/m이 GOPATH로 쓰이는 것 같다.
// 그리고 그 나머지 부분 src/unit39/calc 을 적어주면 되는 듯.
// 이상없이 작동함.

func main() {
	fmt.Println(calc.Sum(2,3))
}
