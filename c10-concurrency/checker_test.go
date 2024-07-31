package c10_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"example.com/c10-concurrency"
)

func TestCheckWebsites(t *testing.T) {
	// Arrange
	websites := []string{
		"https://google.com",
		"https://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	// Act
	actual := c10.CheckWebsites(mockWebsiteChecker, websites)

	// Assert
	expected := map[string]bool{
		"https://google.com":          true,
		"https://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":     false,
	}
	assert.Equal(t, expected, actual)
}

func BenchmarkCheckWebsites(b *testing.B) {
	// Arrange
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()

	// Act
	for i := 0; i < b.N; i++ {
		c10.CheckWebsites(slowWebsiteChecker, urls)
	}
}

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
