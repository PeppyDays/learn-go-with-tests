package c01_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c01-hello-world"
)

func TestHello(t *testing.T) {
	t.Run("shows the default greeting message", func(t *testing.T) {
		// Act
		actual := c01.Hello("", "")

		// Assert
		expected := "Hello, World"
		assert.Equal(t, expected, actual)
	})

	t.Run("shows the targeted greeting message", func(t *testing.T) {
		// Arrange
		name := "Arine"

		// Act
		actual := c01.Hello(name, "")

		// Assert
		expected := "Hello, Arine"
		assert.Equal(t, expected, actual)
	})

	t.Run("shows the greeting message in Spanish", func(t *testing.T) {
		// Arrange
		name := "Elodie"
		language := "Spanish"

		// Act
		actual := c01.Hello(name, language)

		// Assert
		expected := "Hola, Elodie"
		assert.Equal(t, expected, actual)
	})

	t.Run("shows the greeting message in french", func(t *testing.T) {
		// Arrange
		name := "Madam"
		language := "French"

		// Act
		actual := c01.Hello(name, language)

		// Assert
		expected := "Bonjour, Madam"
		assert.Equal(t, expected, actual)
	})
}
