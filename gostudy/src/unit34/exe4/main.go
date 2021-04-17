package main

import (
	"fmt"
)

func main() {
	c:= make(chan int)

	go func() {
		for i:=0; i<5; i++ {
			c<-i
		}
		close(c)
	}()

	for i:= range c {
		fmt.Println(i)
	}
 }

 // range 와 close 함수 특징
 // 이미 닫힌 채널에 값을 보내면 패닉 발생
 // 채널을 닫으면 range 루프가 종료
 // 채널이 열려 있고, 값이 들어오지 않으면 range는 
 // 실행되지 않고 계속 대기. 채널에 값이 들어오면, 그 때부터
 // range가 계속 반복됨