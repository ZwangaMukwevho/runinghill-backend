package resource

import (
	"net/http"

	"github.com/ZwangaMukwevho/backend-api/pkg/model"
	"github.com/gin-gonic/gin"
)

func getWords(c *gin.Context) {

	var words = []model.Word{
		{ID: "1", Word: "Yes", Type: "Noun"},
		{ID: "1", Word: "Yes", Type: "Noun"},
		{ID: "1", Word: "Yes", Type: "Noun"},
	}
	
    c.IndentedJSON(http.StatusOK, words)
}

func postWord(c *gin.Context) {
	var words = []model.Word{
		{ID: "1", Word: "Yes", Type: "Noun"},
	}

	c.IndentedJSON(http.StatusAccepted,words)
}

func getSentence(c *gin.Context) {
	var sentence = []model.Sentence{
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
	}

	c.IndentedJSON(http.StatusOK, sentence)
}

func postSentence(c *gin.Context) {
	var sentence = []model.Sentence{
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
	}
	c.IndentedJSON(http.StatusAccepted,sentence)
}