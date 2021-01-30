package loader

import "github.com/alextanhongpin/go-graphql-template/domain/user"

type DataLoader struct {
	User UserLoader
}

func NewDataLoader(usersvc user.Service) *DataLoader {
	return &DataLoader{
		User: NewUser(usersvc),
	}
}
