package resource

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/db"
	"github.com/ZwangaMukwevho/backend-api/pkg/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	FirebaseClient *db.Client
}

func (h *Handler) getWords(c *gin.Context) {

	var words map[string]model.Word

	ref := h.FirebaseClient.NewRef("words")
	if err := ref.Get(context.Background(), &words); err != nil {
		log.Panic(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	wordsArr := make([]model.Word, 0, len(words))
	for _, word := range words {
		wordsArr = append(wordsArr, word)
	}

	c.IndentedJSON(http.StatusOK, wordsArr)
}

func (h *Handler) postWords(c *gin.Context) {
	var errArray []string
	var words = []model.Word{}

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&words); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	ref := h.FirebaseClient.NewRef("words")
	for _, word := range words {
		wordRef := ref.Child(word.ID)

		// Set the value in Firebase Realtime Database
		if err := wordRef.Set(context.Background(), word); err != nil {
			errArray = append(errArray, fmt.Sprintf("error uploading sentence with Id %s, with error: %s", word.ID, err))
		}
	}

	if len(errArray) != 0 {
		c.IndentedJSON(http.StatusInternalServerError, errArray)
		return
	}

	c.IndentedJSON(http.StatusOK, words)
}

func (h *Handler) getSentence(c *gin.Context) {
	var sentences map[string]model.Sentence

	ref := h.FirebaseClient.NewRef("sentences")
	if err := ref.Get(context.Background(), &sentences); err != nil {
		log.Panic(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	sentenceArr := make([]model.Sentence, 0, len(sentences))
	for _, sentence := range sentences {
		sentenceArr = append(sentenceArr, sentence)
	}

	c.IndentedJSON(http.StatusOK, sentenceArr)
}

func (h *Handler) postSentence(c *gin.Context) {
	var errArray []string
	var sentences = []model.Sentence{}

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&sentences); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	ref := h.FirebaseClient.NewRef("sentences")
	for _, sentence := range sentences {
		sentenceRef := ref.Child(sentence.ID)

		// Set the value in Firebase Realtime Database
		if err := sentenceRef.Set(context.Background(), sentence); err != nil {
			errArray = append(errArray, fmt.Sprintf("error uploading sentence with Id %s, with error: %s", sentence.ID, err))
		}
	}

	if len(errArray) != 0 {
		log.Panic(errArray)
		c.IndentedJSON(http.StatusInternalServerError, errArray)
		return
	}

	c.IndentedJSON(http.StatusOK, sentences)
}
