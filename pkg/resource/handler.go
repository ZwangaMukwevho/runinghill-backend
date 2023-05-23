package resource

import (
	"net/http"

	"firebase.google.com/go/db"
	"github.com/ZwangaMukwevho/backend-api/pkg/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	FirebaseClient *db.Client
}

func (h *Handler) getWords(c *gin.Context) {

	var words = []model.Word{
		{ID: "1", Word: "Yes", Type: "Noun"},
		{ID: "1", Word: "Yes", Type: "Noun"},
		{ID: "1", Word: "Yes", Type: "Noun"},
	}

	c.IndentedJSON(http.StatusOK, words)
}

func (h *Handler) postWord(c *gin.Context) {
	var words = []model.Word{
		{ID: "1", Word: "Yes", Type: "Noun"},
	}

	c.IndentedJSON(http.StatusAccepted, words)
}

func (h *Handler) getSentence(c *gin.Context) {
	var sentence = []model.Sentence{
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
	}

	c.IndentedJSON(http.StatusOK, sentence)
}

func (h *Handler) postSentence(c *gin.Context) {
	var sentence = []model.Sentence{
		{ID: "1", Sentence: "Yes, I do", UserId: "xxx-1111"},
	}
	c.IndentedJSON(http.StatusAccepted, sentence)
}
