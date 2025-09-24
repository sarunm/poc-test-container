package stores

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductStore_Create_Success(t *testing.T) {
	s := SetupMockStore(t)
	name := "newproduct"
	product, err := s.Products().Create(&Product{Name: name, Price: 20})
	assert.NoError(t, err)
	assert.Equal(t, 1, int(product.ID))
	assert.Equal(t, name, product.Name)
	assert.Equal(t, 20, product.Price)
}

func TestProductStore_Create_Error(t *testing.T) {
	s := SetupMockStore(t)
	// สร้าง product ที่มี ID ซ้ำกันเพื่อทดสอบ error
	_, err := s.Products().Create(&Product{ID: 1, Name: "product1", Price: 20})
	assert.NoError(t, err)

	_, err = s.Products().Create(&Product{ID: 1, Name: "product2", Price: 20})
	assert.NotNil(t, err)
}
