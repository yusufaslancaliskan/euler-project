package euler_project

func findLargestPrimeFactorOfNumber(number int) int {
	i := 2
	for i < number {
		if number%i == 0 {
			number = number / i
		}
		i++
	}
	return i
}
