package algorithm

func Palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	} else if s[0] != s[len(s)-1] {
		return false
	} else {
		return Palindrome(s[1 : len(s)-1])
	}
}
