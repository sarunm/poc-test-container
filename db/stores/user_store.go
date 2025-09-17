package stores

import (
	"gorm.io/gorm"
)

type userStore struct {
	storeBase
}

type UserStore interface {
	Create(user *User) (*User, error)
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{storeBase{db: db}}
}

func (u userStore) Create(user *User) (*User, error) {
	result := &User{}
	query := u.db.Create(user).Find(&result)
	if query.Error != nil {
		return nil, query.Error
	}
	return result, nil
}
