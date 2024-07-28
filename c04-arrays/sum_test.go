package c04_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c04-arrays"
)

func TestSum(t *testing.T) {
	t.Run("sut sums all numbers from slice correctly", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 2, 3}

		// Act
		actual := c04.Sum(numbers)

		// Assert
		expected := 6
		assert.Equal(t, expected, actual)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sut sums each of two slices and generate the result of slice correctly", func(t *testing.T) {
		// Arrange
		numbersA := []int{1, 2}
		numbersB := []int{0, 9}

		// Act
		actual := c04.SumAll(numbersA, numbersB)

		// Assert
		expected := []int{3, 9}
		assert.Equal(t, expected, actual)
	})

	t.Run("sut sums one slice and generate the result of slice correctly", func(t *testing.T) {
		// Arrange
		numbers := []int{1, 1, 1}

		// Act
		actual := c04.SumAll(numbers)

		// Assert
		expected := []int{3}
		assert.Equal(t, expected, actual)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sut gathers all tails of slices correctly", func(t *testing.T) {
		// Arrange
		numbersA := []int{1, 2, 3}
		numbersB := []int{0, 9}

		// Act
		actual := c04.SumAllTails(numbersA, numbersB)

		// Assert
		expected := []int{5, 9}
		assert.Equal(t, expected, actual)
	})

	t.Run("sut gathers 0 if slice is empty", func(t *testing.T) {
		// Arrange
		numbers := []int{}

		// Act
		actual := c04.SumAllTails(numbers)

		// Assert
		expected := []int{0}
		assert.Equal(t, expected, actual)
	})
}
