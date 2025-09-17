package stores

import (
	"gorm.io/gorm"
)

type productStore struct {
	storeBase
}

type ProductStore interface {
	Create(product *Product) (*Product, error)
	List() (*Product, error)
}

func NewProductStore(db *gorm.DB) ProductStore {
	return &productStore{storeBase{db: db}}
}

func (p productStore) Create(product *Product) (*Product, error) {
	result := &Product{}
	query := p.db.Create(product).Find(&result)
	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}

func (p productStore) List() (*Product, error) {

	result := &Product{}

	if err := p.db.Model(&Product{}).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
