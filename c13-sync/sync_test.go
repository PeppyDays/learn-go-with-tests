package c13_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c13-sync"
)

func TestCounter(t *testing.T) {
	t.Run("sut increases the counter 3 times", func(t *testing.T) {
		// Arrange
		sut := c13.NewCounter()

		// Act
		sut.Inc()
		sut.Inc()
		sut.Inc()

		// Assert
		actual := sut.Value()
		expected := 3
		assert.Equal(t, expected, actual)
	})

	t.Run("sut increases the counter safely concurrently", func(t *testing.T) {
		// Arrange
		n := 1000
		sut := c13.NewCounter()

		// Act
		var wg sync.WaitGroup
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				sut.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		// Assert
		assert.Equal(t, n, sut.Value())
	})
}
