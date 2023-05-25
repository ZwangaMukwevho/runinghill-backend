package service

import (
	"github.com/ZwangaMukwevho/backend-api/pkg/enums"
	"github.com/ZwangaMukwevho/backend-api/pkg/model"
)

func GetWordsByType(words map[string]model.Word, wordType enums.WordType) []model.Word {
	wordsTypeArr := make([]model.Word, 0, len(words))
	for _, word := range words {
		if word.Type == string(wordType) {
			wordsTypeArr = append(wordsTypeArr, word)
		}
	}
	return wordsTypeArr
}
