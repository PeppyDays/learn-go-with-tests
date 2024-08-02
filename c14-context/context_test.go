package c14_test

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"example.com/c14-context"
)

func TestServer(t *testing.T) {
	t.Run("sut writes fetched response from store correctly", func(t *testing.T) {
		// Arrange
		data := "hello, world!"
		sut := c14.Server(NewSpyStore(data))
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		// Act
		sut.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, data, response.Body.String())
	})

	t.Run("sut tells store to cancel work if request is cancelled", func(t *testing.T) {
		// Arrange
		data := "hello, world!"
		store := NewSpyStore(data)
		sut := c14.Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(ctx)
		response := NewSpyResponseWriter()

		// Act
		sut.ServeHTTP(response, request)

		// Assert
		assert.False(t, response.written)
	})
}

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func NewSpyStore(response string) *SpyStore {
	return &SpyStore{response: response}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func NewSpyResponseWriter() *SpyResponseWriter {
	return &SpyResponseWriter{written: false}
}
