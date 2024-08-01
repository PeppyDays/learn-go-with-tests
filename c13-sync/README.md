# Sync

- Mutex
  - Allows to add locks on data
  - Traditional for sync, communicate by sharing memory
- Channel
  - Share memory by communicating
  - Think it as moving ownership
- When to use what?
  - Use channels when passing ownership of data
  - Use mutexes for managing state

