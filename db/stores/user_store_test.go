package stores

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_UserStore_Create_Success(t *testing.T) {
	s := SetupMockStore(t)
	name := "newuser"
	user, err := s.Users().Create(&User{Name: name})
	assert.Nil(t, err)

	assert.Equal(t, 1, int(user.ID))
	assert.Equal(t, name, user.Name)
}

func Test_UserStore_Create_Error(t *testing.T) {
	s := SetupMockStore(t)
	// สร้าง user ที่มี ID ซ้ำกันเพื่อทดสอบ error
	_, err := s.Users().Create(&User{ID: 1, Name: "user1"})
	assert.Nil(t, err)

	_, err = s.Users().Create(&User{ID: 1, Name: "user2"})
	assert.NotNil(t, err)
}
