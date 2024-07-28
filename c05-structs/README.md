# Structs, Methods and Interfaces

## Interface

### Definition

- A variable of interface type can store a value of any type that is in the type set of the interface
- Basic interface specifies list of methods
  - The name of each explicity specified method must be unique and not blank
- Embedded interface embeds two or more interfaces like `ReadWriter`

### Characteristic

- Interface allows to make functions that can be used with different types
  - This characteristic is called [parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)
- Decouples implementation details
- Still maintain type-safety
- Interface resolution is implicit, so `XXX implements interface YYY` is not required

## Table-driven tests

Summarised from [Go Wiki: Table Driven Tests](https://go.dev/wiki/TableDrivenTests).

- Given a table of test cases, the actual test simply iterates through all table entries
- Basically test of go runs in sequence so that add `t.Parallel()` in the test function if needed
  - When running test, `-parallel n` should be specified to run tests in parallel
