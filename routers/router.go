package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sarunm/poc-test-container/db/stores"
	"github.com/sarunm/poc-test-container/handlers"
	"net/http"
)

func NewRoutes(store stores.StoreBase) http.Handler {
	r := gin.Default()

	productRoutes(r, store)
	userRoutes(r, store)

	return r
}

func userRoutes(r *gin.Engine, store stores.StoreBase) gin.IRoutes {
	user := handlers.NewUserHandler(store)

	u := r.Group("/user")
	{
		u.POST("", user.Create)
	}

	return r
}

func productRoutes(r *gin.Engine, store stores.StoreBase) gin.IRoutes {
	product := handlers.NewProductHandler(store)

	p := r.Group("/product")
	{
		p.POST("", product.Create)
	}

	return r
}
