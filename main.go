package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := NewTestHttp()
	server.Test()
	http.ListenAndServe(":8080", server)
}

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return mux
}

func GinServer() http.Handler {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	//return &router

	return router
}

type testHttp struct {
}

func NewTestHttp() *testHttp {
	return &testHttp{}
}
func (t *testHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (t *testHttp) Test() string {
	return "test"
}
