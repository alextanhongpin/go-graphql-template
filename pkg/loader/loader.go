package loader

import "github.com/alextanhongpin/go-graphql-template/domain/entity"

type DataLoader struct {
	User UserLoader
}

func NewDataLoader(repo entity.Querier) *DataLoader {
	return &DataLoader{
		User: NewUser(repo),
	}
}
