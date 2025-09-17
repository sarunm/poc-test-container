package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sarunm/poc-test-container/db/stores"
)

type UserHandler struct {
	store stores.StoreBase
}

func NewUserHandler(store stores.StoreBase) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) Create(c *gin.Context) {

	var user = &stores.User{}

	user, err := h.store.Users().Create(user)
	if err != nil {
		panic(err)
	}

	c.JSON(200, user)

}
