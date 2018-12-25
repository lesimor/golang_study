package slice

import "fmt"

func Example_insert() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(IntSliceInsert1(a, 3, 25))
	fmt.Println(IntSliceInsert2(a, 3, 25))
	// Output:
	// [10 20 30 25 40 50]
	// [10 20 30 25 40 50]
}

func Example_delete() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(IntSliceDelete(a, 2))
	// Output:
	// [10 20 40 50]
}

func Example_deleteWithoutOrder() {
	// 순서가 중요하지 않은 경우.
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(IntSliceDeleteWithoutOrder(a, 2))
	// Output:
	// [10 20 50 40]
}

func Example_deleteMultiple() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Index 3부터 5개를 삭제.
	fmt.Println(IntSliceDeleteMultiple(a, 3, 5))
	// Output:
	// .
}
