package deduplicate_slice

import "strings"

func UniqueNames(a []string, b []string) []string {
	// Try to come up with your own solution!

	//combined := map[string]string{}
	//AppendToCombined(a, combined)
	//AppendToCombined(b, combined)

	var result []string

	//for _, v := range combined {
	//	result = append(result, v)
	//}

	return result
}

func AppendToCombined(l []string, res map[string]string) {
	for _, i := range l {
		if res[i] == "" {
			res[strings.ToLower(i)] = i
		}
	}
}
