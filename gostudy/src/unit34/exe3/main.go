package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)

	count:= 4

	go func() {
		for i:=0; i<count; i++ {
			done<-true
			fmt.Println("go routine : ", i)
			time.Sleep(1*time.Microsecond)
		}
	}()

	for i:=0; i<count; i++ {
		<-done
		fmt.Println("main func : ", i)

	}
}