package c03_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c03-for"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a returns aaaaa", func(t *testing.T) {
		// Act
		actual := c03.Repeat("a")

		// Assert
		expected := "aaaaa"
		assert.Equal(t, expected, actual)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c03.Repeat("a")
	}
}
