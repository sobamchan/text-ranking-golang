package libs

import (
	"math"
	"strings"
)

func getKeys(m map[string]float64) []string {
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}

type TFIDF struct {
	Documents []string
	DF        map[string]float64
	Vocab     []string
}

func TFIDFNew(documents []string) TFIDF {
	tfidf := TFIDF{documents, map[string]float64{}, []string{}}
	tfidf.CountDF()
	tfidf.Vocab = getKeys(tfidf.DF)
	return tfidf
}

func (tfidf *TFIDF) CountDF() {
	var seen map[string]int
	for _, document := range tfidf.Documents {
		seen = make(map[string]int)
		for _, word := range strings.Fields(document) {
			if _, ok := seen[word]; ok {
				continue
			}
			seen[word] = 1
			if _, ok := tfidf.DF[word]; ok {
				tfidf.DF[word] += 1
			} else {
				tfidf.DF[word] = 1
			}
		}
	}
}

func CalcTF(s string) map[string]float64 {
	tf := make(map[string]float64)
	words := strings.Fields(s)
	for _, word := range words {
		if _, ok := tf[word]; ok {
			tf[word]++
		} else {
			tf[word] = 1.0
		}
	}

	for word, count := range tf {
		tf[word] = count / float64(len(words))
	}
	return tf
}

type TfidfResult []float64
type TfidfResults []TfidfResult

func (tfidf TFIDF) TFIDF(s string) TfidfResult {
	tf := CalcTF(s)
	result := make(TfidfResult, len(tfidf.Vocab))
	vocab := tfidf.Vocab
	for i, w := range vocab {
		idf := math.Log(float64(len(tfidf.Documents))/tfidf.DF[w]) + 1.0
		result[i] = tf[w] * idf
	}
	return result
}

func (tfidf TFIDF) Fit() TfidfResults {
	tfidfs := make(TfidfResults, len(tfidf.Documents))
	for i, document := range tfidf.Documents {
		tfidfs[i] = tfidf.TFIDF(document)
	}
	return tfidfs
}
