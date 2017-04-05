package quantum

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var regexSpace = regexp.MustCompile(`([\.,;\?])`)
var regexVowel = regexp.MustCompile(`(^|\W)([Aa]) ([aeiou])`)
var regexRemoveSpaces = regexp.MustCompile(` ([,\.;\?])`)
var regexPrefix = regexp.MustCompile(`- `)
var regexAddSpace = regexp.MustCompile(`\?(\w)`)
var regexAddSentencesSpace = regexp.MustCompile(`([\.\?])(\w)`)

func Generate(numberOfSentences int) (fulltext string) {
	sentencesCopy := make([][]string, len(sentences))
	copy(sentencesCopy, sentences)

	for i := 0; i < numberOfSentences; i++ {
		fulltext += generateSentence(sentencesCopy[rand.Intn(len(sentencesCopy))])

		if i < (numberOfSentences - 1) {
			fulltext += " "
		}
	}

	return
}

func generateSentence(sentenceTopic []string) (sentence string) {
	patternNumber := rand.Intn(len(sentenceTopic))
	pattern := sentenceTopic[patternNumber]

	pattern = regexSpace.ReplaceAllString(pattern, " $1")
	patternSplit := strings.Split(pattern, " ")

	for _, word := range patternSplit {
		if bullshit, ok := words[word]; ok {
			word = bullshit[rand.Intn(len(bullshit))]
		}

		sentence += fmt.Sprintf("%s ", word)
	}

	sentence = regexVowel.ReplaceAllString(sentence, "$1$2n $3")
	sentence = strings.TrimSpace(sentence)
	sentence = fmt.Sprintf("%s%s", strings.ToUpper(sentence[0:1]), sentence[1:])

	sentence = regexRemoveSpaces.ReplaceAllString(sentence, "$1")
	sentence = regexPrefix.ReplaceAllString(sentence, "-")
	sentence = regexAddSpace.ReplaceAllString(sentence, "? $1")

	return
}
