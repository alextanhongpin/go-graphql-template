package query

// We can place args in a folder called args.

type Query struct {
	*AccountQuery
	*UserQuery
}

func New() *Query {
	return &Query{
		AccountQuery: NewAccountQuery(),
		UserQuery:    NewUserQuery(),
	}
}
