package c11

import (
	"errors"
	"net/http"
	"time"
)

func Racer(timeout time.Duration, urls ...string) (string, error) {
	if len(urls) == 0 {
		return "", errors.New("urls cannot be empty")
	}
	ch := make(chan string, 1)
	for i := 0; i < len(urls); i++ {
		ping(urls[i], ch)
	}
	select {
	case url := <-ch:
		return url, nil
	case <-time.After(timeout):
		return "", errors.New("timed out waiting for responses from URLs")
	}
}

func ping(url string, ch chan string) {
	go func() {
		_, _ = http.Get(url)
		ch <- url
	}()
}
