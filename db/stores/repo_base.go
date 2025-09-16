package stores

import (
	"context"
	"gorm.io/gorm"
)

type RepoBase interface {
	WithTransaction(ctx context.Context, fn func(r RepoBase) error) error
	Users() UserStore
	Products() ProductStore
}

type baseStore struct {
	db *gorm.DB
}
