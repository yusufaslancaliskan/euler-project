package euler_project

func findSumOfEvenFibonacciNumbersLessThan(limit int) int {
	var i, b, sum int = 2, 1, 0
	for i < limit {
		if i%2 == 0 {
			sum += i
		}
		i = i + b
		b = i - b
	}
	return sum
}
