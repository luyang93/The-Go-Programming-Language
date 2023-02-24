package strings

// Join concatenates the elements of array to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(sep string, strs ...string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	n := len(sep) * (len(strs) - 1)
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}

	b := make([]byte, n)
	bp := copy(b, strs[0])
	for _, s := range strs[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
