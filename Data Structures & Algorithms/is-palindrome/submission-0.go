func isPalindrome(s string) bool {
    s = strings.ToLower(s)

	i, j := 0, len(s)-1

	for i < j {

		for i < j && !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
		}

		for i < j && !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
