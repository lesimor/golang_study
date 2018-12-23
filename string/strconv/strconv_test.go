package strconv

import (
	"fmt"
	"strconv"
)

func Example_strconv() {

	i, _ := strconv.Atoi("350")
	k, _ := strconv.ParseInt("cc7fdd", 16, 32)
	s := strconv.Itoa(340)

	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(s)

	// Output:
	// 350
	// 13402077
	// 340
}
