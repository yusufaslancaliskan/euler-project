package euler_project

func findLargestPalindromeMadeFromTheProductOfTwoThreeDigitsNumber() int {
	i, j := 999, 999
	palindromeNumber := 0
	for i > 0 {
		for j > 0 {
			product := i * j
			if isPalindrome(product) {
				if product > palindromeNumber {
					palindromeNumber = product
				}
			}
			j--
		}
		i--
		j = 999
	}
	return palindromeNumber
}

func isPalindrome(product int) bool {
	s := make([]int, 0, 10)
	for product > 0 {
		if product/10 > 0 {
			s = append(s, product%10)
			product = product / 10
		} else {
			s = append(s, product)
			product = 0
		}
	}
	left := s[0 : len(s)/2]
	right := s[len(s)/2 : len(s)]

	for i := 0; i < len(s)/2; i++ {
		if left[i] != right[len(right)-i-1] {
			return false
		}
	}
	return true
}
