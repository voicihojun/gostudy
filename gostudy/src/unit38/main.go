package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

// 책에는 github.com/golang/example/stringutil을 import
// 하라고 나와 있지만 안됨.
// https://github.com/golang/example 사이트에 방문하면
// go get golang.org/x/example/hello 을 실행하라고 나와서 실행한후
// import github.com/golang/example/stringutil 실행하니
// module declares its path as: golang.org/x/example
// but was required as: github.com/golang/example   메세지 나옴
// import golang.org/x/example/stringutil 으로 바꾸니 실행됨.
func main() {
	fmt.Println(stringutil.Reverse("hello"))
}

