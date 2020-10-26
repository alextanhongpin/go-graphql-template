## Test

## Resolver
- test all nested resolvers
- test if the ctx is passed in, and each resolver should be able to resolve
- test if the parent entity is passed down
- test if user context is correct
- test for nil errors
- test for interface and union types resolvers

## Mutation
- happy path
- empty input (one field, two fields, all fields)
- bad input (invalid, non-empty value)
- no auth, with auth
- different owner (can't edit things that don't belong to you)
- test returned resolvers (see above)

## Query
- same as resolver test
- empty input (one field, two fields, all fields)
- bad input (invalid, non-empty value)
- no auth, with auth
- different owner (can't view things that don't belong to you)
- test returned resolvers
- test filter/sort/pagination (default and provided)
