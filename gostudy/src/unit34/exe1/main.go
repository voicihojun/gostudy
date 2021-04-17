package main

import "fmt"

func sum(a,b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int)

	go sum(1, 2, c) //channel을 매개변수로 받는 함수는 
	// 항상 go routine 사용

	n := <-c // <-c 는 채널에 값이 들어올때가지 대기하다가 
	// 값이 들어오면 n 에 값을 할당하고 다음 코드를 실행
	fmt.Println(n)
}