package loader

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/graphql-server-starter/entity"
	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"
)

type UserLoader interface {
	Load(ctx context.Context, id string) (entity.User, error)
	LoadMany(ctx context.Context, ids []string) ([]entity.User, []error)
	Set(ctx context.Context, q entity.User) dataloader.Interface
}

type userFinder interface {
	FindUsersWithIDs(ctx context.Context, ids []uuid.UUID) ([]entity.User, error)
}

type User struct {
	loader *dataloader.Loader
}

func NewUser(finder userFinder) *User {
	l := new(User)

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var ids []uuid.UUID
		for _, key := range keys {
			id, err := uuid.Parse(key.String())
			if err != nil {
				ids = append(ids, id)
			}
		}

		results := make([]*dataloader.Result, len(ids))
		users, err := finder.FindUsersWithIDs(ctx, ids)
		if err != nil {
			for i := range ids {
				results[i] = &dataloader.Result{
					Error: fmt.Errorf("error finding user with id: %s", ids[i]),
				}
			}
			return results
		}

		for i := range users {
			results[i] = &dataloader.Result{
				Data:  users[i],
				Error: nil,
			}
		}
		return results
	}

	// TODO: See if it's possible to use redis.
	cache := &dataloader.NoCache{}
	l.loader = dataloader.NewBatchedLoader(batchFn, dataloader.WithCache(cache))
	return l
}

func (u *User) Load(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	result, err := u.loader.Load(ctx, dataloader.StringKey(id))()
	if err != nil {
		return user, err
	}
	return result.(entity.User), nil
}

func (u *User) LoadMany(ctx context.Context, ids []string) ([]entity.User, []error) {
	keys := make([]dataloader.Key, len(ids))
	for i, id := range ids {
		keys[i] = dataloader.StringKey(id)
	}
	result, err := u.loader.LoadMany(ctx, keys)()
	if err != nil {
		return nil, err
	}
	users := make([]entity.User, len(result))
	for i, r := range result {
		users[i] = r.(entity.User)
	}
	return users, nil
}

// Set adds the loaded entity to the cache.
func (u *User) Set(ctx context.Context, user entity.User) dataloader.Interface {
	key := dataloader.StringKey(user.ID.String())
	return u.loader.Prime(ctx, key, user)
}
