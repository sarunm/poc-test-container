package main

import (
	"context"
	"github.com/sarunm/poc-test-container/db/stores"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	store := stores.NewSqlStore(db)

	users, err := store.Users().Create(&stores.User{Name: "sarun"})
	if err != nil {
		panic(err)
	}
	println(users.Name)

	txn := store.WithTransaction(context.Background(), func(r stores.RepoBase) error {

		users, err := r.Users().Create(&stores.User{Name: "sarun"})
		if err != nil {
			panic(err)
		}
		println(users.Name)

		products, err := r.Products().Create(&stores.Product{Name: "product1", Price: 100})
		if err != nil {
			panic(err)
		}
		println(products.Name)

		return nil

	})

	if txn != nil {
		panic(txn)
	}

	// server := NewSer
}

func NewServer() {

}
