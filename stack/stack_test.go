package stack

import "fmt"

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("3 + 4 * 5"))
	fmt.Println(Eval("2 * ( 3 + 2 )"))
	fmt.Println(Eval("2 - 1"))
	// Output:
	// 5
	// 23
	// 10
	// 1
}
