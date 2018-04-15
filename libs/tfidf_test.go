package libs

import (
	"fmt"
	"testing"
)

func TestTFIDF(t *testing.T) {
	documents := []string{
		"i am a dog .",
		"i like cats with a cat .",
		"i am the boss .",
	}
	tfidf := TFIDFNew(documents)
	results := tfidf.Fit()
	fmt.Println(tfidf.Vocab)
	fmt.Println(tfidf.DF)
	for _, result := range results {
		fmt.Println(result)
	}
}
