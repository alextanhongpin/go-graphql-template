package mutation

type Mutation struct {
	*AccountMutation
	*UserMutation
}

func New() *Mutation {
	return &Mutation{
		AccountMutation: NewAccountMutation(),
		UserMutation:    NewUserMutation(),
	}
}
