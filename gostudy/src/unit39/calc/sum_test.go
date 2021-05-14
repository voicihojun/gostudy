package calc

import "testing"

// ========= Test Code 1 ============
//
// func TestSum(t *testing.T) {
// 	r := Sum(1, 2)
// 	if r != 3 {
// 		t.Error("결과값이 3이 아닙니다. r = ", r)
// 	}
// }

// ========= Test Code 2 ============
//
// type TestData struct {
// 	argument1 int
// 	argument2 int
// 	result int
// }

// var testdata = []TestData {
// 	{2, 6, 8},
// 	{-8, 3, -5},
// 	{6, -6, 0},
// 	{0, 0, 1},
// }

// func TestSum(t *testing.T) {
// 	for _, d := range testdata {
// 		r := Sum(d.argument1, d.argument2)
// 		if r != d.result {
// 			t.Errorf(
// 				"%d + %d의 결과값이 %d가 아닙니다 r=%d",
// 				// error string에는 마침표(.)를 쓰면 안됨.
// 				// 에러메시지 : "Errorf format %  is missing verb at end of string"
// 				d.argument1,
// 				d.argument2,
// 				d.result,
// 				r,

// 			)
// 		}
// 	}
// }

// ========= Test Code 3 ============
// benchmark test : 성능 측정

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1,2)
	}
}
//Benchmark로 함수이름 시작. 그 다음 글자는 꼭 대문자. *testing.B를 매개변수로 받음
