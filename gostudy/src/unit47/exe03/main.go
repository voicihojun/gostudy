package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s []byte = make([]byte, 0)

	s = strconv.AppendBool(s, true)
	fmt.Println(string(s))

	s = strconv.AppendFloat(s, 1.3, 'f', -1, 32)
	fmt.Println(string(s))

	s = strconv.AppendInt(s, -10, 10)
	fmt.Println(string(s))

	s = strconv.AppendUint(s, 32, 16)
	fmt.Println(string(s))


}