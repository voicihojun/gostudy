package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")

	fmt.Println(r.Replace("<div><span>Hello, world!</span></div>"))
}