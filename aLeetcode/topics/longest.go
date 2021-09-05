package main

func longestPara(s string) (int, string) {

	res := ""
	length := 0

	for i := range s {
		if len(s) - i < length{
			return length, res
		}
		for j := i; j < len(s); j++ {
			// +1 means includeing the last one
			if issymmetry(s[i : j+1]) {
				if j-i +1 > length {
					length = j - i + 1
					res = s[i : j+1]
				}
			}
		}
	}

	return length, res
}

func issymmetry(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
