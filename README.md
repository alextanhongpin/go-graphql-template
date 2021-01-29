# Graphql golang

- [ ] authentication
- [ ] rate limiting
- [ ] designing graphql schema
- [ ] dataloaders

- [x] integrate background worker
- [x] integrate redis to calculate page views
- [x] get client ip + id for unique page views
- [] notification logic
- [] answer/question count logic
- [] comments logic


## Tests Assertions
- using goconvey and testify suite/assertions

https://github.com/smartystreets/goconvey/wiki/Assertions


## Pre-commit
To enable, change the `.git/hooks/pre-commit`:
```
#!/bin/sh

sh scripts/pre-commit.sh
[ $? -ne 0 ] && exit 1;
exit 0;
```
