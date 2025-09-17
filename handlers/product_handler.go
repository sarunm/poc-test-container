package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sarunm/poc-test-container/db/stores"
)

type ProductHandler struct {
	store stores.StoreBase
}

func NewProductHandler(store stores.StoreBase) *ProductHandler {
	return &ProductHandler{
		store: store,
	}
}

type CreateProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (c *CreateProductRequest) toProductModel() *stores.Product {
	return &stores.Product{
		Name:  c.Name,
		Price: c.Price,
	}
}

func (h *ProductHandler) Create(c *gin.Context) {

	var req CreateProductRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		panic(err)
	}

	createdProduct, err := h.store.Products().Create(req.toProductModel())
	if err != nil {
		panic(err)
	}

	c.JSON(200, createdProduct)
}

func (h *ProductHandler) List(c *gin.Context) {

	response, err := h.store.Products().List()
	if err != nil {
		panic(err)
	}

	c.JSON(200, response)
}
