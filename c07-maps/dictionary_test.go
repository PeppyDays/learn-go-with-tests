package c07_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c07-maps"
)

func TestSearch(t *testing.T) {
	// Arrange
	dictionary := c07.Dictionary{"test": "this is just a test"}

	t.Run("sut find known word correctly", func(t *testing.T) {
		// Act
		actual, _ := dictionary.Search("test")

		// Assert
		expected := "this is just a test"
		assert.Equal(t, expected, actual)
	})

	t.Run("sut throws error for searching unknown word", func(t *testing.T) {
		// Act
		_, err := dictionary.Search("hello")

		// Assert
		assert.Error(t, err)
	})
}

func TestAdd(t *testing.T) {
	// Arrange
	dictionary := c07.Dictionary{}

	t.Run("sut adds word not existing in the dictionary correctly", func(t *testing.T) {
		// Act
		_ = dictionary.Add("hello", "world")

		// Assert
		actual, _ := dictionary.Search("hello")
		expected := "world"
		assert.Equal(t, expected, actual)
	})

	t.Run("sut updates word existing in the dictionary correctly", func(t *testing.T) {
		// Act
		_ = dictionary.Add("hello", "world")
		err := dictionary.Add("hello", "world")

		// Assert
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	// Arrange
	dictionary := c07.Dictionary{
		"hello": "world",
		"test":  "this is just a test",
	}

	t.Run("sut updates word existing in the dictionary correctly", func(t *testing.T) {
		// Act
		err := dictionary.Update("hello", "there")

		// Assert
		assert.NoError(t, err)
		actual, _ := dictionary.Search("hello")
		assert.Equal(t, "there", actual)
	})

	t.Run("sut throws error for updating unknown word", func(t *testing.T) {
		// Act
		err := dictionary.Update("arine", "ok?")

		// Assert
		assert.Error(t, err)
	})
}
