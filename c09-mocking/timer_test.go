package c09_test

import (
	"bytes"
	"testing"

	"example.com/c09-mocking"
	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	t.Run("sut prints 3, 2, 1 and Go! correctly", func(t *testing.T) {
		// Arrange
		buffer := &bytes.Buffer{}
		dummySleeper := &DoubleSleeper{}

		// Act
		c09.Countdown(buffer, dummySleeper)

		// Assert
		actual := buffer.String()
		expected := "3\n2\n1\nGo!"
		assert.Equal(t, expected, actual)
	})

	t.Run("sut calls sleeper correctly", func(t *testing.T) {
		// Arrange
		buffer := &bytes.Buffer{}
		spySleeper := &DoubleSleeper{}

		// Act
		c09.Countdown(buffer, spySleeper)

		// Assert
		actual := spySleeper.calls
		expected := 3
		assert.Equal(t, expected, actual)
	})

	t.Run("sut calls writer and sleeper in a row correctly", func(t *testing.T) {
		// Arrange
		spyOperationRecorder := &DoubleOperationRecorder{}

		// Act
		c09.Countdown(spyOperationRecorder, spyOperationRecorder)

		// Assert
		actual := spyOperationRecorder.operations
		expected := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}
		assert.Equal(t, expected, actual)
	})
}

type DoubleSleeper struct {
	calls int
}

func (s *DoubleSleeper) Sleep() {
	s.calls++
}

type DoubleOperationRecorder struct {
	operations []string
}

func (r *DoubleOperationRecorder) Write(p []byte) (n int, err error) {
	r.operations = append(r.operations, "write")
	return 0, nil
}

func (r *DoubleOperationRecorder) Sleep() {
	r.operations = append(r.operations, "sleep")
}
