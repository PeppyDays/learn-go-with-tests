package c02_test

import (
	"fmt"
	"testing"

	"example.com/c02-integers"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("adder with 2 and 2 returns 4", func(t *testing.T) {
		// Act
		actual := c02.Add(2, 2)

		// Assert
		expected := 4
		assert.Equal(t, expected, actual)
	})
}

func ExampleAdd() {
	sum := c02.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
