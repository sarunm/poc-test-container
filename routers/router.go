package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sarunm/poc-test-container/db/stores"
	"github.com/sarunm/poc-test-container/handlers"
	"github.com/sarunm/poc-test-container/middlewares"
	"log/slog"
	"net/http"
	"os"
)

func NewRoutes(store stores.StoreBase) http.Handler {
	r := gin.Default()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	api := r.Group("/api", middlewares.GlobalExceptionHandlerAndLogger(logger))
	productRoutes(api, store)
	userRoutes(api, store)

	return r
}

func userRoutes(r *gin.RouterGroup, store stores.StoreBase) gin.IRoutes {
	user := handlers.NewUserHandler(store)

	u := r.Group("/user")
	{
		u.POST("", user.Create)
	}

	return r
}

func productRoutes(r *gin.RouterGroup, store stores.StoreBase) gin.IRoutes {
	product := handlers.NewProductHandler(store)

	p := r.Group("/product")
	{
		p.POST("", product.Create)
	}

	return r
}
