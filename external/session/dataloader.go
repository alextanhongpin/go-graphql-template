package session

import (
	"context"

	"github.com/alextanhongpin/go-graphql-template/pkg/loader"
	"github.com/alextanhongpin/pkg/contextkey"
)

var dataLoaderKey = contextkey.Key("dataloader")

func DataLoader(ctx context.Context) *loader.DataLoader {
	dl, _ := dataLoaderKey.Value(ctx).(*loader.DataLoader)
	return dl
}

func WithDataLoader(ctx context.Context, dl *loader.DataLoader) context.Context {
	return context.WithValue(ctx, dataLoaderKey, dl)
}
