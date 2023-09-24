package main

import (
	"strconv"
)

func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	toStr := strconv.Itoa(x)
	strLen := len(toStr)
	strBytes := make([]byte, strLen)

	for i := 0; i < strLen; i++ {
		strBytes[i] = toStr[strLen-1-i]
	}

	return string(strBytes) == toStr
}
