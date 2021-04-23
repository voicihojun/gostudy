package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "HI RURURURU"
	r := strings.NewReader(s)
	fmt.Println(r)

	var s1, s2 string
	n, _ := fmt.Fscanf(r, "%s %s", &s1, &s2)
	fmt.Println("입력개수: ", n)
	fmt.Println(s1)
	fmt.Println(s2)

}