package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sarunm/poc-test-container/db/stores"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert" // library ยอดนิยม ช่วยให้เขียน assert ง่ายขึ้น
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Products() stores.ProductStore {
	args := m.Called()
	return args.Get(0).(stores.ProductStore)
}
func (m *MockStore) Users() stores.UserStore {
	args := m.Called()
	return args.Get(0).(stores.UserStore)
}

func (m *MockStore) WithTransaction(ctx context.Context, fn func(r stores.StoreBase) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}

type MockProductStore struct {
	mock.Mock
}

func (m *MockProductStore) Create(product *stores.Product) (*stores.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*stores.Product), args.Error(1)
}

func (m *MockProductStore) List() (*stores.Product, error) {
	args := m.Called()
	return args.Get(0).(*stores.Product), args.Error(1)
}

func TestProductHandler_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockStore := new(MockStore)
	mockProductStore := new(MockProductStore)

	mockStore.On("Products").Return(mockProductStore)
	mockProductStore.On("Create", mock.AnythingOfType("*stores.Product")).Return(&stores.Product{
		ID:    1,
		Name:  "Test Product",
		Price: 100,
	}, nil)

	handler := NewProductHandler(mockStore)

	recorder := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/products", handler.Create)

	requestPayload := gin.H{"name": "Test Product", "price": 100}
	jsonBody, _ := json.Marshal(requestPayload)

	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)

	var response stores.Product
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Product", response.Name)
	assert.Equal(t, 100, response.Price)

	mockProductStore.AssertExpectations(t)
	mockStore.AssertExpectations(t)
}
