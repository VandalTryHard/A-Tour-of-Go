package prose

import "strings"

func JoinWithCommas(pharses []string) string {
	// НЕ проходит join_test.go
	// result := strings.Join(pharses[:len(pharses)-1], ", ")
	// result += ", and "
	// result += pharses[len(pharses)-1]
	// return result

	// Проходит join_test.go
	if len(pharses) == 0 {
		return ""
	} else if len(pharses) == 1 {
		return pharses[0]
	} else if len(pharses) == 2 {
		return pharses[0] + " and " + pharses[1]
	} else {
		result := strings.Join(pharses[:len(pharses)-1], ", ")
		result += ", and "
		result += pharses[len(pharses)-1]
		return result
	}
}
