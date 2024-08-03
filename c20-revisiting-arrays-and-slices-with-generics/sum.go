package c20

func Sum(numbers []int) int {
	return reduce(numbers, func(x int, y int) int { return x + y }, 0)
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

func reduce[A any](collection []A, f func(A, A) A, initial A) A {
	var result = initial
	for i := 0; i < len(collection); i++ {
		result = f(result, collection[i])
	}
	return result
}
