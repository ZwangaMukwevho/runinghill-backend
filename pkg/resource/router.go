package resource

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {

	router := gin.Default()
    router.GET("/albums", getWords)

	return router
}