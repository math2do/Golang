package main

func isPalindrome(num int) bool {
	var inputNum int = num
	res := 0
	for num > 0 {
		rem := num % 10
		res = (res * 10) + rem
		num = num / 10
	}

	if inputNum == res {
		return true
	} else {
		return false
	}
}
