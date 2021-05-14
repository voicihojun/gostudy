package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("../exe01/hello.txt.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	r, err := gzip.NewReader(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	// 결과가 unexpected EOF가 나오는데, 
	// 아마도 cloud에 파일이 올라가 있기 때문인것 같다. 
	// 클라우드에서 다운로드 받은 후 실행하면 정상작동됨. 
}