package stack

import "fmt"

func ExamplePush() {
	s := make([]int, 0, 10)
	Push(s, 10)
	fmt.Println(Push(s, 10))
	// Output:
	// .
}

func ExamplePop() {
	s := []int{1,2,3,4,5}
	var popped int

	s, popped = Pop(s)
	fmt.Println(popped, s)

	s, popped = Pop(s)
	fmt.Println(popped, s)

	s, popped = Pop(s)
	fmt.Println(popped, s)
	// Output:
	// 5 [1 2 3 4]
	// 4 [1 2 3]
	// 3 [1 2]
}

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
