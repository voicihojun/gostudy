// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

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
		os.FileMode(0644),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.WriteString("Hello, gogogo")
	w.Flush()

	r:= bufio.NewReader(file)

	fi, _ := file.Stat()

	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)
	// os.SEEK_SET, os.SEEK_CUR, os.SEEK_END deprecated 되었음.
	// io.SeekStart / io.SeekCurrent / io.SeekEnd

	r.Read(b)
	fmt.Println(string(b))
}