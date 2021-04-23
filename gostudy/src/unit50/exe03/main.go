package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := "ABCDEFGGGGG"
	r := strings.NewReader(s)
	io.Copy(os.Stdout, r)
}