package c11_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"example.com/c11-select"
)

func TestRacer(t *testing.T) {
	t.Run("sut returns error when comparing without urls", func(t *testing.T) {
		// Arrange
		timeout := time.Millisecond

		// Act
		_, err := c11.Racer(timeout)

		// Assert
		assert.Error(t, err)
	})

	t.Run("sut chooses fast server correctly", func(t *testing.T) {
		// Arrange
		timeout := time.Second
		slowServer, slowURL := makeDeplayedServer(20 * time.Millisecond)
		fastServer, fastURL := makeDeplayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		// Act
		actual, _ := c11.Racer(timeout, slowURL, fastURL)

		// Assert
		expected := fastURL
		assert.Equal(t, expected, actual)
	})

	t.Run("sut returns an error if a servers don't respond within specified timeout", func(t *testing.T) {
		// Arrange
		timeout := 10 * time.Millisecond
		serverA, serverAURL := makeDeplayedServer(50 * time.Millisecond)
		serverB, serverBURL := makeDeplayedServer(60 * time.Millisecond)
		defer serverA.Close()
		defer serverB.Close()

		// Act
		_, actual := c11.Racer(timeout, serverAURL, serverBURL)

		// Assert
		assert.Error(t, actual)
	})
}

func makeDeplayedServer(duration time.Duration) (*httptest.Server, string) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
	return server, server.URL
}
