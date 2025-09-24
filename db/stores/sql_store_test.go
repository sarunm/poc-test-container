package stores

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupMockStore(t *testing.T) StoreBase {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	// สร้างตารางที่จำเป็น
	err = db.AutoMigrate(&User{})
	err = db.AutoMigrate(&Product{})
	assert.NoError(t, err)

	return NewSqlStore(db)
}

func Test_WithTransaction_Success(t *testing.T) {

	store := setupMockStore(t)
	err := store.WithTransaction(nil, func(r StoreBase) error {
		user, err := r.Users().Create(&User{Name: "testuser"})
		assert.NoError(t, err)
		assert.Equal(t, "testuser", user.Name)
		return nil
	})
	assert.NoError(t, err)

	db := store.(*sqlStore).db
	var foundUser User
	result := db.First(&foundUser, "name = ?", "testuser")
	assert.NoError(t, result.Error)
	assert.Equal(t, "testuser", foundUser.Name)
}

func Test_WithTransaction_Fail(t *testing.T) {
	store := setupMockStore(t)
	err := store.WithTransaction(nil, func(r StoreBase) error {
		_, err := r.Users().Create(&User{Name: "testuser"})
		assert.NoError(t, err)
		return assert.AnError // จำลองความผิดพลาดเพื่อให้เกิดการ rollback
	})
	assert.Error(t, err)

	db := store.(*sqlStore).db
	var foundUser User
	result := db.First(&foundUser, "name = ?", "testuser")
	assert.Error(t, result.Error) // ควรจะไม่พบ user เนื่องจาก rollback
	assert.Equal(t, int64(0), result.RowsAffected)
}
