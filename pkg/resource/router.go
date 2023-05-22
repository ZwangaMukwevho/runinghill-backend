package resource

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {

	router := gin.Default()
    router.GET("/word", getWords)
	router.POST("/word",postWord)
	router.GET("/sentence",getSentence)
	router.POST("/sentence",postSentence)

	return router
}