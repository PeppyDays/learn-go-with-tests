package c04

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func SumAll(listOfNumbers ...[]int) []int {
	lengthOfNumbers := len(listOfNumbers)
	sums := make([]int, lengthOfNumbers)
	for i := 0; i < lengthOfNumbers; i++ {
		sums[i] = Sum(listOfNumbers[i])
	}
	return sums
}

func SumAllTails(listOfNumbers ...[]int) []int {
	lengthOfNumbers := len(listOfNumbers)
	sums := make([]int, lengthOfNumbers)
	for i := 0; i < lengthOfNumbers; i++ {
		numbers := listOfNumbers[i]
		if len(numbers) <= 1 {
			sums[i] = 0
		} else {
			numbersInTail := numbers[1:]
			sums[i] = Sum(numbersInTail)
		}
	}
	return sums
}
