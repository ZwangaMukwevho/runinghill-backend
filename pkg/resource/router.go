package resource

import "github.com/gin-gonic/gin"

func NewRouter(handler Handler) *gin.Engine {

	router := gin.Default()
	router.GET("/words", handler.getWords)
	router.POST("/words", handler.postWords)
	router.GET("/words/:type", handler.getWordsByType)
	router.GET("/sentences", handler.getSentence)
	router.POST("/sentences", handler.postSentence)

	return router
}
