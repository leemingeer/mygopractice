package main

func longestPalindrome(s string) string {
	res  := string(s[0])
	length := 1

	for i:=0;i<len(s);i++{
		for j:=i;j<len(s);j++{
			if isPalindrome(s[i:j+1]){
				if j - i +1 > length{
					length = j - i +1
					res = s[i:j+1]
				}
			}
		}
	}
	return res
}

func  isPalindrome(s string) bool{
	for i,j := 0,len(s) - 1; i < j; i,j = i +1, j-1{
		if s[i] != s[j]{
			return false
		}
	}
	return true
}