package main

import (
	"bytes"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/nimrodoron/restserver/pkg/storage"
	"net/http"
)

func main() {
	db := storage.CreateInMemoryStorage()
	router := gin.Default()
	initRouter(router, db)
	pprof.Register(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initRouter(router *gin.Engine, storage datasource) {

	router.GET("/resource/:name", gin.BasicAuth(gin.Accounts{ "nimrod.oron@sap.com": "12345678"}),func(c *gin.Context) {
		resource, err := storage.Retrieve(c.Param("name"))
		if err == nil {
			c.JSON(http.StatusOK, resource)
		} else {
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	})

	router.PUT("/resource/:name", gin.BasicAuth(gin.Accounts{ "nimrod.oron@sap.com": "12345678"}), func(c *gin.Context) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		resource := buf.String()
		storage.Persist(c.Param("name"), resource)
	})

	router.DELETE("/resource/:name", gin.BasicAuth(gin.Accounts{ "nimrod.oron@sap.com": "12345678"}), func(c *gin.Context) {
		storage.Delete(c.Param("name"))
	})

	router.GET("/resources", gin.BasicAuth(gin.Accounts{ "nimrod.oron@sap.com": "12345678"}), func(c *gin.Context) {
		resource, _ := storage.RetrieveAll()
		if len(resource) > 0 {
			c.JSON(http.StatusOK, resource)
		} else {
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	})
}

type datasource interface {
	Persist(name, content string) error
	Retrieve(name string) (string, error)
	RetrieveAll() ([]string, error)
	Delete(name string) error
}
