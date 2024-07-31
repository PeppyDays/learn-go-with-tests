# Select

## HTTP Test

`net/http/httptest` enables to easily create a mock HTTP server.

```go
import "net/http/httptest"

server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  time.Sleep(time.Second)
  w.WriteHeader(http.StatusOK)
}))
fmt.Println(server.URL)
```

This helps slow and flaky test originated from HTTP server. Furthermore, some edge cases (e.g. testing handling of error from third-party HTTP API) can be testable.

## Select

- `chan struct{}` is the smallest data type from a memory perspective
- `select` chooses a fastest one from multiple channels
- `case <-time.After(10 * time.Second)` can be used for timeout

Some other patterns or tips for channel and `select` are found here.

- https://hamait.tistory.com/1017
- https://velog.io/@moonyoung/golang-channel-with-select-%ED%97%B7%EA%B0%88%EB%A6%AC%EB%8A%94-%EC%BC%80%EC%9D%B4%EC%8A%A4-%EC%A0%95%EB%A6%AC%ED%95%98%EA%B8%B0
