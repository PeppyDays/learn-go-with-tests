package c05_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c05-structs"
)

func TestPerimeter(t *testing.T) {
	t.Run("sut calculates rectangle's perimeter correctly", func(t *testing.T) {
		// Arrange
		width, height := 10.0, 10.0
		rectangle := c05.Rectangle{width, height}

		// Act
		actual := rectangle.Perimeter()

		// Assert
		expected := 40.0
		assert.Equal(t, expected, actual)
	})
}

func TestArea(t *testing.T) {
	areaTestCases := []struct {
		shape    c05.Shape
		expected float64
	}{
		{shape: &c05.Rectangle{12, 6}, expected: 72.0},
		{shape: &c05.Circle{10.0}, expected: 314.1592653589793},
		{shape: &c05.Triangle{12.0, 6.0}, expected: 36.0},
	}

	for _, testCase := range areaTestCases {
		t.Run(
			fmt.Sprintf("%v area is calculated correctly", testCase.shape),
			func(t *testing.T) {
				t.Parallel()
				actual := testCase.shape.Area()
				assert.Equal(t, testCase.expected, actual)
			},
		)
	}
}
