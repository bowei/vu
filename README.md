# vu

Log file viewer

# todo

- parsing will yield a tree of nodes.
- nodes are denoted by open, close patterns.
- depending on the given rules, seeing a non-nestable open will implicitly 
  result in a close being emitted (a la HTML DOM error recovery parsing)
- fuzzy parsing (e.g. regexes) for common things such as time stamps, errors etc.
- add in classification and summarization logic?