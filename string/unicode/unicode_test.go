package unicode

import "fmt"

func Example_print_unicode() {
	// 문자열은 읽기 전용.
	s := "가각갂"
	for i, r := range s {
		fmt.Println(i, r)
	}
	fmt.Println()
	// Output:
	// 0 44032
	// 3 44033
	// 6 44034
}

func Example_print_string_with_unicode() {
	var (
		start rune = 44032 // 가
		end   rune = 44059 // 갛
	)
	for i := start; i <= end; i++ {
		fmt.Println(string(i))
	}
	// Output:
	// 가
	// 각
	// 갂
	// 갃
	// 간
	// 갅
	// 갆
	// 갇
	// 갈
	// 갉
	// 갊
	// 갋
	// 갌
	// 갍
	// 갎
	// 갏
	// 감
	// 갑
	// 값
	// 갓
	// 갔
	// 강
	// 갖
	// 갗
	// 갘
	// 같
	// 갚
	// 갛
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}
