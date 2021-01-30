package loader

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"

	"github.com/alextanhongpin/go-graphql-template/domain/user"
)

type UserLoader interface {
	Load(ctx context.Context, id string) (user.User, error)
	LoadMany(ctx context.Context, ids []string) ([]user.User, []error)
	Set(ctx context.Context, q user.User) dataloader.Interface
}

type userFinder interface {
	FindUsersWithIDs(ctx context.Context, ids []uuid.UUID) ([]user.User, error)
}

type User struct {
	loader *dataloader.Loader
}

func NewUser(finder userFinder) *User {
	l := new(User)

	// TODO: Check implementation.
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

func (u *User) Load(ctx context.Context, id string) (user.User, error) {
	var usr user.User
	result, err := u.loader.Load(ctx, dataloader.StringKey(id))()
	if err != nil {
		return usr, err
	}
	return result.(user.User), nil
}

func (u *User) LoadMany(ctx context.Context, ids []string) ([]user.User, []error) {
	keys := make([]dataloader.Key, len(ids))
	for i, id := range ids {
		keys[i] = dataloader.StringKey(id)
	}
	result, err := u.loader.LoadMany(ctx, keys)()
	if err != nil {
		return nil, err
	}
	users := make([]user.User, len(result))
	for i, r := range result {
		users[i] = r.(user.User)
	}
	return users, nil
}

// Set adds the loaded entity to the cache.
func (u *User) Set(ctx context.Context, usr user.User) dataloader.Interface {
	key := dataloader.StringKey(usr.ID.String())
	return u.loader.Prime(ctx, key, usr)
}
