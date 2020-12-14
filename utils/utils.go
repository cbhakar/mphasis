package utils

import "strconv"

func StrInListStatus(s string, l []string) (r bool) {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}

func CheckIntValue(input string) (int, bool) {
	v, err := strconv.Atoi(input)
	if err != nil {
		return v, false
	}
	return v, true
}
