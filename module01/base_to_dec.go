package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14

func BaseToDec(value string, base int) int {
	abc := "0123456789ABCDEF"
	cba := make(map[rune]int)
	for i, v := range abc {
		cba[v] = i
	}
	ret := 0
	for _, v := range value {
		ret = ret*base + cba[v]
	}
	return ret
}
