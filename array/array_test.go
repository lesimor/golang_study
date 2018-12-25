package array

import (
	"fmt"
	"github.com/lesimor/golang_study/string/hangul"
)

func Example_array() {
	// 갯수를 미리 정해놓으면 array type
	fruits := [3]string{"사과", "바나나", "레몬"}
	var postPosition string
	for _, fruit := range fruits {
		// 은/는 처리
		if postPosition = "는"; hangul.HasConsonantSuffix(fruit) {
			postPosition = "은"
		}
		fmt.Printf("[range]%s%s 맛있다\n", fruit, postPosition)
	}

	// for i := range fruits {} <= lhs가 하나인 경우에는 index를 의미.

	for i := 0; i < len(fruits); i++ {
		if postPosition = "는"; hangul.HasConsonantSuffix(fruits[i]) {
			postPosition = "은"
		}
		fmt.Printf("[len]%s%s 맛있다\n", fruits[i], postPosition)
	}
	// Output:
	// [range]사과는 맛있다
	// [range]바나나는 맛있다
	// [range]레몬은 맛있다
	// [len]사과는 맛있다
	// [len]바나나는 맛있다
	// [len]레몬은 맛있다
}
