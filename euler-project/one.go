package euler_project


func findMultiplesOfThreeAndFive(r int) int {
	sum:=0;
	for i := 0; i < r; i++ {
		if i % 3 == 0 || i % 5 == 0  {
			sum += i;
		}
	}
	return sum
}