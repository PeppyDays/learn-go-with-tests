package c08_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/c08-dependency-injection"
)

func TestGreet(t *testing.T) {
	// Arrange
	buffer := bytes.Buffer{}

	// Act
	c08.Greet(&buffer, "Chris")

	// Assert
	actual := buffer.String()
	expected := "Hello, Chris"
	assert.Equal(t, expected, actual)
}
