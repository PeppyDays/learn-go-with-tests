# Maps

- Map is not a reference, it is a pointer to a `runtime.hmap` structure
- Map can be nil, so the initialization is mandatory like `make(map[string]int)` or `m := map[string]int{}`
- Map is not safe for concurrent
  - If needed, use `sync.RWMutex` together
