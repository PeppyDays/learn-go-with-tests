# Pointers and Errors

## Pointer

- All parameters and structs defined in functions and methods are referenced by value, so that copy is done implicitly
- To minimize copy over memory or mutate an object, pointer should be used
- Pointers can be nil, so that nil check or initialisation (with `make`) should be done in the first place
- Struct pointers provides automatic dereferencing of the member attributes
  - `(*w).balance` could be `w.balance`

## New Type

- New type can be defined with `type` command
- New type helps adding domain specific meaning to values
- New type can have methods

## Error

- Checking for a string in an error would result in a flaky test
  - Never inspect the output of error because it is not usually the concern of the test
- Modify an naive error into meaningful value would help to avoid checking an error message

## Handling an Error in Go

Summarised from [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully).

### Sentinel Errors

```go
var ErrSomething = errors.New("oh?")
if err == ErrSomething { .. }
```

- Sentinel errors become part of the public API
- Sentinel errors create tight coupling between two packages
  - The caller must compare the result to predeclared value
  - If the error signiture has changed, the caller and callee should change something
- Sentinel errors doesn't have context-aware messages
- Avoid sentinel errors

### Error Types

```go
type ErrSomething struct {
  Message string
  Line int
}

func (e *ErrSomething) Error() string {
  return fmt.Sprintf("%v: %v", e.Line, e.Message)
}

err := something()
switch err := err.(type) {
case nil:
  // call succeeded
case *ErrSomething:
  fmt.Println("error occurred on line:", err.Line)
default:
  // unknown error
}
```

- Error has ability to wrap an underlying error to provide more context
- Error type must be public that means still strongly coupled
- Error type must be asserted to divide the reaction of the errors like written above
- Avoid error types

### Opaque Errors

- Use an error as the sign of success or fail, and does not meaning into the message
- The most flexible error handle strategy as it does not require coupling between two parties
- If binary approach (ok or not) is not sufficient, assert errors for behaviour, not type
  - For instance, implement `temporary` interface on retriable errors, and check retriability from the caller if error is changeable into `temporary`
- For error lineage, use error wrapping with `fmt.Errorf("...: %v", err)`
  - In the caller side, check an error is propagated from error X by using `errors.Is(err, X)`
