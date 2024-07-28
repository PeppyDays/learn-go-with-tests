# Hello, World

## Rules of Writing Tests

- Test file needs to be in a file with a name like `xxx_test.go`
- Module of test files are `xxx_test` if possible

## TDD Discipline

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
  - Check if this test produces an easy to understand description of the failure
  - Check if this test is relevant for our requirements
    - If this test passes without writing any code, codes in somewhere does something more that we don't know
- Write enough code to pass the test
- Refactor
  - Helps ensure to design good software by refactoring


