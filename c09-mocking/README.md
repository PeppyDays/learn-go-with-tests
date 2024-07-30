# Mocking

## Mocking is Evil?

If mocking code is becoming complicated or having too much mocks to test something, you should *listen* to the bad feeling and think about the code. It is a sign of:

- The code under the test does too many things
- Its dependencies are too fine-grained
- Test is concerned with implementation details

Make TDD strengthful as a signal of bad design. If you're having troubles with tests and mocks, consider these:

- Refactoring doesn't have to modify tests if tests are focusing on behaviour, not implementation
- If a test uses more than 3 mocks, then it is a red flag

Without mocking, it is difficult to test under these situations:

- Calling a external service that should fail?
- Making a system in a particular state?
- Calling something slow?
- ..

Finally, focus on what it does, not the way it works. Tests have the value on the behaviour test and future refactoring.

## Test Double

Summarised from [Test Double](https://martinfowler.com/bliki/TestDouble.html) and [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html).

- Dummy
  - Passed around but never actually used meaningfully
- Fake
  - Have working implementation, but not for production
  - Verify state
- Stub
  - Provide canned answer to calls
  - Verify state
- Spy
  - Stub + record some information based on how they were called
  - Verify state
- Mock
  - Have pre-programmed with expectations with form a specification of the calls they are expected to receive
  - Verify behaviour
