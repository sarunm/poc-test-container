package stores

import (
	"gorm.io/gorm"
)

type userStore struct {
	baseStore
}

type UserStore interface {
	Create(user *User) (*User, error)
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{baseStore{db: db}}
}

func (u userStore) Create(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}
