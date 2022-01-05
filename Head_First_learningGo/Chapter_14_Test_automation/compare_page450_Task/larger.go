package compare_page450_Task

func Larger(a int, b int) int {
	// НЕ проходит larger_test.go
	// if a < b {
	// 	return a
	// } else {
	// 	return b
	// }

	// Проходит larger_test.go
	if a > b {
		return a
	} else {
		return b
	}
}
