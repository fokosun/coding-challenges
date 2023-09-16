package string_manipulation

func ReverseStr(s string) string {
	strLen := len(s)
	strBytes := make([]byte, strLen)

	for i := 0; i < strLen; i++ {
		strBytes[i] = s[strLen-1-i]
	}

	return string(strBytes)
}

// ReverseSliceOfStrings reverses the items in a slice of strings without reversing the strings themselves
func ReverseSliceOfStrings(s []string) []string {
	strLen := len(s)
	values := make([]string, strLen)

	for i := 0; i < strLen; i++ {
		values[i] = s[strLen-1-i]
	}

	return values
}
