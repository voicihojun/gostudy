package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	w := bufio.NewWriter(file)

	fmt.Fprintf(w, "%d, %f, %s", 1, 1.1, "Hello")
	w.Flush()

	r := bufio.NewReader(file)

	fi, _ := file.Stat()

	b := make([]byte, fi.Size())
	file.Seek(0, io.SeekStart)
	r.Read(b)
	fmt.Println(string(b))
}