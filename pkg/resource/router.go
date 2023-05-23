package resource

import "github.com/gin-gonic/gin"

func NewRouter(handler Handler) *gin.Engine {

	router := gin.Default()
	router.GET("/word", handler.getWords)
	router.POST("/word", handler.postWords)
	router.GET("/sentence", handler.getSentence)
	router.POST("/sentence", handler.postSentence)

	return router
}
