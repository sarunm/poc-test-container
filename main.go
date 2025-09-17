package main

import (
	"github.com/sarunm/poc-test-container/db/stores"
	"github.com/sarunm/poc-test-container/routers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	db := NewSqlServer()

	store := stores.NewSqlStore(db)
	router := routers.NewRoutes(store)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// อย่าลืมจัดการเรื่อง Graceful Shutdown ในโค้ดจริง
	server.ListenAndServe()
}

func NewSqlServer() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&stores.User{}, &stores.Product{}); err != nil {
		panic(err)
	}
	//
	//store := stores.NewSqlStore(db)
	//
	//users, err := store.Users().Create(&stores.User{Name: "sarun"})
	//if err != nil {
	//	panic(err)
	//}
	//println(users.Name)
	//
	//txn := store.WithTransaction(context.Background(), func(r stores.StoreBase) error {
	//
	//	users, err := r.Users().Create(&stores.User{Name: "sarun"})
	//	if err != nil {
	//		panic(err)
	//	}
	//	println(users.Name)
	//
	//	products, err := r.Products().Create(&stores.Product{Name: "product1", Price: 100})
	//	if err != nil {
	//		panic(err)
	//	}
	//	println(products.Name)
	//
	//	return nil
	//
	//})
	//
	//if txn != nil {
	//	panic(txn)
	//}
	return db
}
