package slice

import "fmt"

func Example_slicing() {
	// 대괄호에 숫자를 넣지 않으면 slice가 된다.
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	// Output:
	// .
}

// 슬라이스 덧붙이기.
func Example_append() {
	f1 := []string{"사과", "딸기", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
	// .
}

// 슬라이스 capacity
func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()

	// sliced1으로 해도 뒤에 남은 공간이 그대로 있으므로 참조 가능.
	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	// 초과시 두배씩 capacity가 늘어남.
	extendedNums := append(nums, 6)
	fmt.Println(extendedNums)
	fmt.Println("len:", len(extendedNums))
	fmt.Println("cap:", cap(extendedNums))
	fmt.Println()

	// 할당하지 않은 값에 대해서는??
	makeSlice := make([]int, 5)
	makeSlice[0] = 100
	fmt.Println(makeSlice[:4]) // 따로 설정되지 않은 경우 0
	// Output:
	// .

}

func Example_sliceCopy() {
	// 안정적으로 copy하는 방법.
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
}

func Example_copyMethod() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, 3)

	// copy는 복사된 element의 갯수를 반환.
	// min(src 길이, dest의 길이)
	if n := copy(dest, src); n != len(src) {
		fmt.Println("복사가 덜 됐습니다.")
	}

	// Output:
	// .
}
