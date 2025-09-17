package stores

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type sqlStore struct {
	db           *gorm.DB
	userStore    UserStore
	productStore ProductStore
}

func NewSqlStore(db *gorm.DB) StoreBase {
	return &sqlStore{db: db,
		userStore:    NewUserStore(db),
		productStore: NewProductStore(db)}
}

func (s sqlStore) Users() UserStore {
	return s.userStore
}

func (s sqlStore) Products() ProductStore {
	return s.productStore
}

func (s sqlStore) WithTransaction(ctx context.Context, fn func(r StoreBase) error) error {
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer tx.Rollback()

	txStore := &sqlStore{
		db:           tx,
		userStore:    NewUserStore(tx),
		productStore: NewProductStore(tx),
	}

	err := fn(txStore)
	if err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit().Error
}
