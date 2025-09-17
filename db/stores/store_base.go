package stores

import (
	"context"
	"gorm.io/gorm"
)

type StoreBase interface {
	WithTransaction(ctx context.Context, fn func(r StoreBase) error) error
	Users() UserStore
	Products() ProductStore
}

type storeBase struct {
	db *gorm.DB
}
