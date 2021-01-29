package resolver

// Resolver represents the root for all queries and mutations.
type Resolver struct {
	*AccountQuery
	*AccountMutation
	*UserQuery
	*UserMutation
}

// New returns a Resolver configured with the Options.
func New() *Resolver {
	return &Resolver{
		AccountQuery:    NewAccountQuery(),
		AccountMutation: NewAccountMutation(),
		UserQuery:       NewUserQuery(),
		UserMutation:    NewUserMutation(),
	}
}
