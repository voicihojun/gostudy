package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string = strconv.FormatBool(true)
	fmt.Println(s1)

	var s2 string = strconv.FormatFloat(1.3, 'e', -1, 32)
	// strconv.FormatFloat(f, fmt, prec, bitsize)
	// f: 변환할 실수값s1
	// fmt : 어떤 형식으로 나타낼 건지...s1
	// - 'b' / 'e' : 1.23e+07 / 'E' : 1.23E+07 / 'f' : 1.342
	// prec : 실수의 정밀도(지수를 제외한 숫자의 자리수). -1을 지정하면 적절?하게 바꿔줌.s1
	// bitSize: 부동소수점 비트 수. 32 or 64
	fmt.Println(s2)

	var s3 string = strconv.FormatInt(-10, 2) // -10을 10진수로된 문자열로 변환
	fmt.Println(s3)

	var s4 string = strconv.FormatUint(32, 16) // -10을 10진수로된 문자열로 변환
	fmt.Println(s4)


}