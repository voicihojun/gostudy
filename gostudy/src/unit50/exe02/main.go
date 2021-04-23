package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt", 
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	s:= "Hi hihihihihi hohohoho"
	r := strings.NewReader(s)

	w := bufio.NewWriter(file)

	// w.ReadFrom(r) 아래로 대체해도 됨. 
	io.Copy(w, r)
	w.Flush()

	rr := bufio.NewReader(file)

	fi, _ := file.Stat()

	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)
	rr.Read(b)
	fmt.Println(string(b))





}