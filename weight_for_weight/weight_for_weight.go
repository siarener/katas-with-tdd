package kata

import (
	"sort"
	"strconv"
	"strings"
)

type weight struct {
	word     string
	digitSum int
}

func sumDigitsFromString(word string) int {
	var charSum int
	for _, char := range word {
		digit, err := strconv.Atoi(string(char))
		if err == nil {
			charSum += int(digit)
		}
	}
	return charSum
}

func OrderWeight(strng string) string {
	var allWeights []weight
	realWeight := strings.Split(strng, " ")
	for _, word := range realWeight {
		weightDigitSum := sumDigitsFromString(word)
		allWeights = append(allWeights, weight{
			word:     word,
			digitSum: weightDigitSum,
		})
	}
	sort.Slice(allWeights, func(i, j int) bool {
		if allWeights[i].digitSum != allWeights[j].digitSum {
			return allWeights[i].digitSum < allWeights[j].digitSum
		}
		return allWeights[i].word < allWeights[j].word
	})

	orderedWords := extractWeights(allWeights)
	return strings.Join(orderedWords, " ")
}

func extractWeights(allWeights []weight) []string {
	var orderedWeightWords []string
	for _, w := range allWeights {
		orderedWeightWords = append(orderedWeightWords, w.word)
	}
	return orderedWeightWords
}
