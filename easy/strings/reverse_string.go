package reverse_strings

func ReverseStr(s string) string {
	strLen := len(s)
	strBytes := make([]byte, strLen)

	for i := 0; i < strLen; i++ {
		strBytes[i] = s[strLen-1-i]
	}

	return string(strBytes)
}
