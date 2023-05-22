package resource

import (
	"net/http"

	"github.com/ZwangaMukwevho/backend-api/pkg/model"
	"github.com/gin-gonic/gin"
)

func getWords(c *gin.Context) {

	var words = []model.Words{
		{ID: "1", Word: "Yes", Type: "Noun"},
		{ID: "1", Word: "Yes", Type: "Noun"},
		{ID: "1", Word: "Yes", Type: "Noun"},
	}
	
    c.IndentedJSON(http.StatusOK, words)
}

func postWords(c *gin.Context) {
	
}