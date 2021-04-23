package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string] interface{})

	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.Marshal(data)
	fmt.Println(string(doc))

	doc, _ = json.MarshalIndent(data, "", "  ")
	// data : json 문서로 만들 데이터
	// 두번째 매개변수 :json 문서의 첫칸에 표시할 문자열(prefix)
	// 세번째 매개변수 : 들여쓰기할 문자
	fmt.Println(string(doc))
}