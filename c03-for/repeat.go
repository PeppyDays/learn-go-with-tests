package c03

import "strings"

const repeatCount = 5

func Repeat(c string) string {
	// strings.Repeat is 2x slower than using strings.Builder
	// return strings.Repeat(c, repeatCount)
	var sb strings.Builder
	for i := 0; i < repeatCount; i++ {
		sb.WriteString("a")
	}
	return sb.String()
}
