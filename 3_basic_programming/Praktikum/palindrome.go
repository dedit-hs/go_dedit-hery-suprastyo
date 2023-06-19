package basicProgramming

func Palindrome(kata string) bool {
	var balik string

	for i := len(kata) - 1; i >= 0; i-- {
		balik += string(kata[i])
	}

	if kata == balik {
		return true
	} else {
		return false
	}
}
