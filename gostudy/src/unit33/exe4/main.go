package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, world!"

	// Hello, world! 100 .....
	// go routine 은 반복문이 완전히 끝난다음에 생성(cpu 하나 사용시)
	// 여러개 사용해도 0-99까지 빠짐없이 출력은 안됨.
	for i:=0; i<100; i++ {
		go func() {
			fmt.Println(s, i)

		}()
	}

	// 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨줘야함. 
	// 매개변수로 넘겨주는 시점에 해당 변수값이 복사되므로 고루틴에 
	// 그 당시의 변수를 사용할 수 있기 때문.
	for i:=0; i<100; i++ {
		go func(n int) {
			fmt.Println(s, n)

		}(i)
	}

	fmt.Scanln()

}