package libs

import (
	"github.com/ikawaha/kagome/tokenizer"
	"strings"
)

type Parse struct {
	Pos       []string
	Tokenizer tokenizer.Tokenizer
	Normalize bool
}

func NewParse(pos []string, normalize bool) Parse {
	newTokenizer := tokenizer.New()
	parse := Parse{pos, newTokenizer, normalize}
	return parse
}

func (parse Parse) Parse(rawSent string) string {
	tokens := parse.Tokenizer.Tokenize(rawSent)
	posToParse := strings.Join(parse.Pos, " ")
	var words []string

	for _, token := range tokens {
		features := token.Features()
		if token.Class == tokenizer.DUMMY {
			continue
		}
		if strings.Contains(posToParse, features[0]) {
			if parse.Normalize {
				words = append(words, features[len(features)-3])
			} else {
				words = append(words, token.Surface)
			}
		}
	}
	return strings.Join(words, " ")
}
