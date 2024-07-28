# Arrays and Slices

## Blank Identifier

Summarised from [The Blank Identifier in Effective Go](https://go.dev/doc/effective_go#blank).

- The blank identifier can be assigned or declared with any value of any type for discarding harmlessly
- The blank identifier represents a placeholder where a variable is required but the actual value is not needed

## Arrays and Slices

Summaries from [Go Slices: usage and internals](https://go.dev/blog/slices-intro).

### Array

- Array type is defined as a length and an element type
- If array length is different, the two types are different (e.g. `[4]int` is differ from `[5]int`)
- Zero values are filled when initialised
- Array is not a pointer of the first element, it is just a value
  - When operating along with array elements, it is copied
  - If don't want to be copied, use a pointer like `*[5]int`

### Slice

- Slice type is defined as a element type and no specified length
- Slice literal is declared just like an array literal, except the element count
  - `letters := []string{"a", "b"}`
  - `s := make([]byte, 5, 5) // s == []byte{0, 0, 0, 0, 0}`
- Slice consists of a pointer to the array, the length of the segment, and its capacity
- Re-slicing does not copy the slice's data, just create a new slice as a pointer
- When growing a slice, if capacity is overflowed, the new array is created, copied, and the new slice is reassigned
  - `a = append(a, 1, 2, 3)`
