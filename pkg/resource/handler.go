package resource

import (
	"context"
	"encoding/json"
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
	// ref := h.FirebaseClient.NewRef("words")
	// var words []model.Word
	// if err := ref.Get(context.Background(), &words); err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, err)
	// 	// return nil, fmt.Errorf("failed to fetch words from Firebase: %v", err)
	// }

	c.IndentedJSON(http.StatusOK, words)
}

func (h *Handler) postWords(c *gin.Context) {
	var words = []model.Word{}

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&words); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	// var ind string
	// for index, word := range words {
	// 	// id
	// 	// := uuid.New().String()
	// 	ind = "one"
	// 	if index == 1 {
	// 		ind = "two"
	// 	}
	// 	refString := fmt.Sprintf("words/%s", ind)
	// 	ref := h.FirebaseClient.NewRef(refString)
	// 	err := ref.Set(context.TODO(), word)
	// 	if err != nil {
	// 		error := fmt.Errorf("an error occured processing word with Id: %s, with exception: %w", word.ID, err)
	// 		log.Fatal(error)
	// 		c.IndentedJSON(http.StatusInternalServerError, error)
	// 		return
	// 	}
	// }

	// for _, word := range words {
	// 	_, err := ref.Push(context.TODO(), word)
	// 	if err != nil {
	// 		error := fmt.Errorf("an error occured processing word with Id: %s, with exception: %w", word.ID, err)
	// 		log.Fatal(error)
	// 		c.IndentedJSON(http.StatusInternalServerError, error)
	// 		return
	// 	}
	// }
	ref := h.FirebaseClient.NewRef("words")
	for _, word := range words {

		wordRef := ref.Child(word.ID)

		// Convert the word struct to JSON
		wordJSON, err := json.Marshal(word)

		fmt.Print(wordJSON)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
			// return fmt.Errorf("failed to marshal word: %v", err)
		}
		fmt.Print(wordRef)
		// Set the value in Firebase Realtime Database
		if err := wordRef.Set(context.Background(), word); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
			// return fmt.Errorf("failed to upload word to Firebase: %v", err)
		}
	}

	c.IndentedJSON(http.StatusOK, words)
}

func (h *Handler) getSentence(c *gin.Context) {
	var sentence []model.Sentence

	ref := h.FirebaseClient.NewRef("sentence/")
	if err := ref.Get(context.TODO(), &sentence); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, sentence)
}

func (h *Handler) postSentence(c *gin.Context) {

	var sentence = []model.Sentence{}
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&sentence); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	ref := h.FirebaseClient.NewRef("sentence/")
	if err := ref.Set(context.TODO(), sentence); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, sentence)
}
