package byte

import "fmt"

func Example_printByByte() {
	s := "가각갂"
	// len 은 바이트 단위로 연산.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}
	fmt.Println()
	// Output:
	// ea
	// b0
	// 95
	// eb
	// b3
	// 91
	// ec
	// 9a
	// b1
}

func Example_printByCharacter(){
	s := "가나다"
	fmt.Printf("% x", s)
	// Output:
	// ea b0 80 eb 82 98 eb 8b a4
}