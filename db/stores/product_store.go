package stores

import (
	"gorm.io/gorm"
)

type productStore struct {
	baseStore
}

type ProductStore interface {
	Create(product *Product) (*Product, error)
}

func NewProductStore(db *gorm.DB) ProductStore {
	return &productStore{baseStore{db: db}}
}

func (p productStore) Create(product *Product) (*Product, error) {
	//TODO implement me
	panic("implement me")
}
