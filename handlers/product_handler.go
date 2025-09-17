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

func (h *ProductHandler) Create(c *gin.Context) {

	product := &stores.Product{}

	createdProduct, err := h.store.Products().Create(product)
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
