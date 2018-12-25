package slice

func IntSliceInsert1(s []int, idx int, val int) []int {
	newSlice := append(s[:idx+1], s[idx:]...)
	newSlice[idx] = val
	return newSlice
}

func IntSliceInsert2(s []int, idx int, val int) []int {
	// idx 조건 검사 로직 추가.
	var newSlice []int
	if idx < len(s) {
		newSlice = append(s[:idx+1], s[idx:]...)
		newSlice[idx] = val
	} else {
		newSlice = append(s, val)
	}
	return newSlice
}

func IntSliceDelete(s []int, idx int) []int {
	return append(s[:idx], s[idx+1:]...)
}

func IntSliceDeleteWithoutOrder(s []int, idx int) []int {
	s[idx] = s[len(s)-1]
	newSlice := s[:len(s)-1]
	return newSlice
}

func IntSliceDeleteMultiple(s []int, removePosition int, removeLength int) []int {
	copyStart := len(s) - removeLength
	// 삭제 범위와 뒤에서 가져오는 범위가 겹칠 경우
	if removePosition+removeLength > copyStart {
		copyStart = removePosition + removeLength
	}
	copy(s[removePosition:removePosition+removeLength], s[copyStart:])
	return s[:len(s)-removeLength]
}
