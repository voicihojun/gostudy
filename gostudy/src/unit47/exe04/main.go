package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error
	var b1 bool
	b1, err = strconv.ParseBool("true")
	fmt.Println(b1, err)

	var num1 float64
	num1, err = strconv.ParseFloat("1.3", 64)
	fmt.Println(num1, err)

	var num2 int64
	num2, err = strconv.ParseInt("-10", 10, 32)
	// 문자열 내 -10이 10진수임을 나타내고 있고 그걸 부동소수점 비트수 32로 나타내는 것
	fmt.Println(num2, err)

	var num3 uint64
	num3, err = strconv.ParseUint("30", 16, 32)
	// 30이 16진수이므로 10진수로는 48이고 그걸 부동소수점 비트수 32로 나타내는 것
	fmt.Println(num3, err)
}